package crackdict

import (
	"github.com/L4ml3da/zaku/util"
	"strings"
)

func DictRead(param string)[]string {

	if len(param) == 0 {
		return nil
	}
	if strings.Contains(param, "file:") {
		idx := strings.Index(param, ":")
		//fmt.Println("get file " + param[idx + 1:])
		filename := param[idx+1:]
		return util.ReadFile(filename)
	} else {
		return strings.Split(param, ",")
	}
	return nil
}
