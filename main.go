/************************************************************
** @Description: PPGo_Job2
** @Author: haodaquan
** @Date:   2018-06-05 22:24
** @Last Modified by:   haodaquan
** @Last Modified time: 2018-06-05 22:24
*************************************************************/
package main

import (
	"fmt" // 标准库中的fmt包来进行控制台输出
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kardianos/service"

	"PPGo_Job/jobs"
	"PPGo_Job/models"
	_ "PPGo_Job/routers"

	"github.com/astaxie/beego"
)

var logger service.Logger

// Program structures.
//
//	Define Start and Stop methods.
type program struct {
	exit chan struct{}
}

func (p *program) Start(s service.Service) error {

	if service.Interactive() {
		logger.Info("Running in Terminal.")

		//初始化数据模型
		var StartTime = time.Now().Unix()
		models.Init(StartTime)
		jobs.InitJobs()

		beego.Run()

	} else {
		logger.Info("Running under service manager.")
	}
	p.exit = make(chan struct{})

	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}

func (p *program) run() {

	// Do work here
	logger.Info("PPGo_Job is running.")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Working dir=", dir)
	os.Chdir(dir)

	//初始化数据模型
	var StartTime = time.Now().Unix()
	models.Init(StartTime)
	jobs.InitJobs()

	beego.Run()
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	beego.BeeApp.Server.Shutdown(nil)

	// Any work in Stop should be quick, usually a few seconds at most.
	logger.Info("PPGO_Job is stopping!")
	close(p.exit)
	return nil
}

// 原来是 func init(), 使用windows service后就不用了。
func init_depreciated() {
	fmt.Println("Init PPGO_Job.")
	//初始化数据模型
	var StartTime = time.Now().Unix()
	models.Init(StartTime)
	jobs.InitJobs()
}

func main() {
	// fmt.Println("Starting PPGO_Job.")

	var name = "PPGo_Job2.8.0"
	var displayName = "PPGo_Job2.8.0"
	var desc = "PPGo_Job2.8.0"

	svcConfig := &service.Config{
		Name:        name,
		DisplayName: displayName,
		Description: desc,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	// 通过以下代码来控制服务的启动和停止
	if len(os.Args) > 1 {

		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
			return
		}
		return
	}

	// 非 Service 模式
	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
