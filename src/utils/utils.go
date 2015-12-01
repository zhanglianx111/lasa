package utils

import(
	"reflect"
	"fmt"
	"strconv"
	//"errors"
	//"encoding/json"
)

func AnalysisJob(job interface{}) (map[string]interface{}) {
	jobData := make(map[string]interface{})

	valueJ := reflect.ValueOf(job).Elem()
	
	jobData["name"] = valueJ.FieldByName("Name").String()
	jobData["color"] = valueJ.FieldByName("Color").String()
	jobData["description"] = valueJ.FieldByName("Description").String()
	jobData["displayname"] = valueJ.FieldByName("DisplayName").String()
	jobData["url"] = valueJ.FieldByName("URL").String()
	jobData["inQueue"] = valueJ.FieldByName("InQueue").Bool()
	jobData["buildable"] = valueJ.FieldByName("Buildable").Bool()
	jobData["concurrentBuild"] = valueJ.FieldByName("ConcurrentBuild").Bool()

	var db []map[string]interface{}
	bds := valueJ.FieldByName("Builds")
	for m := 0; m < bds.Len(); m++ {		
		db = append(db, map[string]interface{}{"number": bds.Index(m).FieldByName("Number").Int(),
												  "url": bds.Index(m).FieldByName("URL").String()})
	}
	jobData["builds"] = db

	lb := valueJ.FieldByName("LastBuild")
	jobData["lastBuild"] = map[string]interface{}{"number": lb.FieldByName("Number").Int(),
												     "url": lb.FieldByName("URL").String()}

	lfb := valueJ.FieldByName("LastFailedBuild")
	jobData["lastFailedBuild"] = map[string]interface{}{"number": lfb.FieldByName("Number").Int(), 
														   "url": lfb.FieldByName("URL").String()}
	
	lsb := valueJ.FieldByName("LastSuccessfulBuild")
	jobData["lastSuccessfulBuild"] = map[string]interface{}{"number": lsb.FieldByName("Number").Int(),
															   "url": lsb.FieldByName("URL").String()}
	
	lstb := valueJ.FieldByName("LastStableBuild")
	jobData["lastStableBuild"] = map[string]interface{}{"number": lstb.FieldByName("Number").Int(),
														   "url": lstb.FieldByName("URL").String()}

	//fmt.Println(jobData)
	return jobData
}

func AnalysisNode(node interface{}) (map[string]interface{}) {
	nodeData := make(map[string]interface{})

	valueN := reflect.ValueOf(node).Elem()
	fmt.Println("valueN:", valueN)
	nodeData["displayName"] = valueN.FieldByName("DisplayName").String()
	//nodeData["executors"] = valueN.FieldByName("Executors").Int()
	nodeData["idle"] = valueN.FieldByName("Idle").Bool()
	nodeData["numExecutors"] = valueN.FieldByName("NumExecutors").Int()
	nodeData["offline"] = valueN.FieldByName("Offline").Bool()

	/*
	jobData["offlineCause"] = valueJ.FieldByName("OfflineCause").Bool()

	var db []map[string]interface{}
	bds := valueJ.FieldByName("Builds")
	for m := 0; m < bds.Len(); m++ {		
		db = append(db, map[string]interface{}{"number": bds.Index(m).FieldByName("Number").Int(), 
												"url": bds.Index(m).FieldByName("URL").String()})
	}
	jobData["builds"] = db

	lb := valueJ.FieldByName("LastBuild")
	jobData["lastBuild"] = map[string]interface{}{"number": lb.FieldByName("Number").Int(),
													  "url": lb.FieldByName("URL").String()}

	lfb := valueJ.FieldByName("LastFailedBuild")
	jobData["lastFailedBuild"] = map[string]interface{}{"number": lfb.FieldByName("Number").Int(), 
															"url": lfb.FieldByName("URL").String()}
	
	lsb := valueJ.FieldByName("LastSuccessfulBuild")
	jobData["lastSuccessfulBuild"] = map[string]interface{}{"number": lsb.FieldByName("Number").Int(), 
															"url": lsb.FieldByName("URL").String()}
	
	lstb := valueJ.FieldByName("LastStableBuild")
	jobData["lastStableBuild"] = map[string]interface{}{"number": lstb.FieldByName("Number").Int(), 
															"url": lstb.FieldByName("URL").String()}

	//fmt.Println(jobData)
	*/
	return nodeData
}

func StrToInt64(s string) int64 {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return -1
	}
	return n
}

