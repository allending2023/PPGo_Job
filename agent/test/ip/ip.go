/************************************************************
** @Description: ip
** @Author: george hao
** @Date:   2019-06-27 09:22
** @Last Modified by:  george hao
** @Last Modified time: 2019-06-27 09:22
*************************************************************/
package main

import (
	"PPGo_Job/libs"

	"github.com/astaxie/beego/logs"
)

func main() {
	logs.Info(libs.PublicIp())

}
