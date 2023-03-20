package repository

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"ginws/config"

	_ "github.com/sijms/go-ora/v2"
)

type LogInfo struct {
	Username   string
	IPAddress  string
	Handler    string
	BodyParams interface{}
	ErrorInfo  *string
}

func SaveLog(d *config.Dependencies, lg *LogInfo) error {
	sqlStatement := `INSERT INTO log_info (log_id, username, ip_address, handler, body_params, error_info)
                     VALUES (log_info_seq.nextval, :1, :2, :3, :4, :5)`

	bodyParamsJSON, err := json.Marshal(lg.BodyParams)
	if err != nil {
		return err
	}

	stmt, err := d.Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.Exec(lg.Username, lg.IPAddress, lg.Handler, string(bodyParamsJSON), lg.ErrorInfo)
	if err != nil {
		return err
	}

	return nil
}

// if needed - XML save below

type XMLNode struct {
	XMLName xml.Name
	Content string    `xml:",chardata"`
	Nodes   []XMLNode `xml:",omitempty"`
}

func jsonToXML(data []byte) (string, error) {
	var jsonData map[string]interface{}
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return "", err
	}

	node := createNode("l", jsonData)
	xmlData, err := xml.MarshalIndent(node, "", "")
	if err != nil {
		return "", err
	}

	return string(xmlData), nil
}

func createNode(key string, value interface{}) XMLNode {
	node := XMLNode{XMLName: xml.Name{Local: key}}

	switch v := value.(type) {
	case map[string]interface{}:
		for k, val := range v {
			child := createNode(k, val)
			node.Nodes = append(node.Nodes, child)
		}
	case []interface{}:
		for _, val := range v {
			child := createNode(key, val)
			node.Nodes = append(node.Nodes, child)
		}
	default:
		if v != "" && v != nil {
			node.Content = fmt.Sprintf("%v", value)
		}
	}

	return node
}

func SaveLogXML(d *config.Dependencies, logInfo *LogInfo) error {

	sqlStatement := `INSERT INTO log_info (log_id, username, ip_address, handler, body_params, error_info)
                     VALUES (log_info_seq.nextval, :1, :2, :3, :4, :5)`

	bodyParamsJSON, err := json.Marshal(logInfo.BodyParams)
	if err != nil {
		return err
	}

	bodyParamsXML, err := jsonToXML(bodyParamsJSON)
	if err != nil {
		return err
	}

	stmt, err := d.Db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	defer func() {
		_ = stmt.Close()
	}()

	_, err = stmt.Exec(logInfo.Username, logInfo.IPAddress, logInfo.Handler, bodyParamsXML, logInfo.ErrorInfo)
	return err

}
