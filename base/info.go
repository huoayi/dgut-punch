package base

import (
	"encoding/json"
	"time"

	"github.com/golang/glog"
)

func SetForm(user UserAccount, execution string) map[string]string {
	data := make(map[string]string)
	data["username"] = user.Username
	data["password"] = user.Password
	data["captcha"] = ""
	data["_eventId"] = "submit"
	data["cllt"] = "userNameLogin"
	data["dllt"] = "generalLogin"
	data["lt"] = ""
	data["execution"] = execution
	return data
}

func SetSubmitInfo(info []byte) (string, []byte) {
	infoData := make(map[string]interface{})
	err := json.Unmarshal(info, &infoData)
	if err != nil {
		glog.Error(err)
		return "", nil
	}
	subInfo := make(map[string]interface{})
	if infoData["user_data"] == nil {
		glog.Errorf("set submit info err, user_data is nil,info:[%+v]", info)
	}
	userData := infoData["user_data"].(map[string]interface{})
	var faculty map[string]int
	faculty["粤台产业科技学院"] = 1
	faculty["经济与管理学院"] = 1
	faculty["法律与社会工作学院"] = 1
	faculty["继续教育学院"] = 1
	faculty["国际学院"] = 1
	faculty["东莞理工学院与法国国立工艺学院"] = 1

	subInfo["submit_time"] = time.Now().Format("2006-01-02")
	subInfo["name"] = userData["name"]
	subInfo["faculty_name"] = userData["faculty_name"]
	subInfo["class_name"] = userData["class_name"]
	subInfo["username"] = userData["username"]
	subInfo["card_number"] = userData["card_number"]
	subInfo["identity_type"] = setDataWithInt(userData["identity_type"])
	subInfo["remark"] = userData["remark"]
	subInfo["tel"] = userData["tel"]
	subInfo["connect_person"] = userData["connect_person"]
	subInfo["connect_tel"] = userData["connect_tel"]
	subInfo["family_address_detail"] = userData["family_address_detail"]
	subInfo["current_country"] = userData["current_country"]
	subInfo["current_province"] = userData["current_province"]
	subInfo["current_city"] = userData["current_city"]
	subInfo["current_district"] = userData["current_district"]
	subInfo["latest_acid_test"] = userData["latest_acid_test"]
	subInfo["now_in_area_level"] = setDataWithInt(userData["now_in_area_level"])
	subInfo["current_in_city"] = setDataWithInt(userData["current_in_city"])
	subInfo["huji_region"] = userData["huji_region"]
	subInfo["family_region"] = userData["family_region"]
	subInfo["jiguan_region"] = userData["jiguan_region"]
	subInfo["current_region"] = userData["current_region"]
	subInfo["huji_region_name"] = userData["huji_region_name"]
	subInfo["family_region_name"] = userData["family_region_name"]
	subInfo["jiguan_region_name"] = userData["jiguan_region_name"]
	subInfo["card_type"] = userData["card_type"]
	subInfo["campus"] = setDataWithInt(userData["campus"])
	subInfo["have_diagnosis"] = setDataWithInt(userData["have_diagnosis"])
	subInfo["have_gone_important_area"] = setDataWithInt(subInfo["have_gone_important_area"])
	subInfo["have_contact_hubei_people"] = setDataWithInt(subInfo["have_contact_hubei_people"])
	subInfo["have_contact_illness_people"] = setDataWithInt(subInfo["have_contact_illness_people"])
	subInfo["have_isolation_in_dg"] = setDataWithInt(subInfo["have_isolation_in_dg"])
	subInfo["is_in_dg"] = setDataWithInt(subInfo["is_in_dg"])
	subInfo["is_new_in_dg"] = setDataWithInt(subInfo["is_new_in_dg"])
	subInfo["have_go_out"] = setDataWithInt(subInfo["have_go_out"])
	subInfo["is_specific_people"] = setDataWithInt(subInfo["is_specific_people"])
	subInfo["health_code_status"] = setDataWithInt(subInfo["health_code_status"])
	subInfo["in_controllerd_area"] = setDataWithInt(subInfo["in_controllerd_area"])
	subInfo["completed_vaccination"] = setDataWithInt(subInfo["completed_vaccination"])
	subInfo["have_stay_area"] = setDataWithInt(subInfo["have_stay_area"])
	subInfo["family_situation"] = setDataWithArray(subInfo["family_situation"])
	subInfo["health_situation"] = 1
	subInfo["body_temperature"] = "36.5"
	subInfo["is_in_school"] = 1
	if _, ok := faculty[subInfo["faculty_name"].(string)]; ok {
		subInfo["is_in_school"] = 2
	}

	//新增选项增加日期 : 2022.11.10
	subInfo["current_community_and_house"] = userData["current_community_and_house"]
	subInfo["current_street"] = userData["current_street"]
	//gps备用选项
	subInfo["gps_address_name"] = userData["gps_address_name"]
	subInfo["gps_city"] = userData["gps_city"]
	subInfo["gps_city_name"] = userData["gps_city_name"]
	subInfo["gps_country"] = userData["gps_country"]
	subInfo["gps_country_name"] = userData["gps_country_name"]
	subInfo["gps_district"] = userData["gps_district"]
	subInfo["gps_district_name"] = userData["gps_district_name"]
	subInfo["gps_province"] = userData["gps_province"]
	subInfo["gps_province_name"] = userData["gps_province_name"]

	ans := make(map[string]interface{})
	ans["data"] = subInfo
	subByte, err := json.Marshal(ans)
	if err != nil {
		glog.Info(err)
		return "", nil
	}
	if subByte == nil {
		glog.Info("subByte is nil")
		return "", nil
	}

	return userData["name"].(string), subByte
}

func setDataWithInt(data interface{}) interface{} {
	if data == nil {
		return 0
	}
	return data
}
func setDataWithArray(data interface{}) interface{} {
	if data == nil {
		return []int{}
	}
	return data
}
