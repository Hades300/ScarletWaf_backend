package common

import (
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"strings"
)

type RegisterForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (this *RegisterForm) Validate() error {
	return validation.ValidateStruct(this,
		validation.Field(&this.Name, validation.Required, validation.Length(3, 20)),
		validation.Field(&this.Email, is.Email.Error("请输入合法的邮箱地址")),
	)
}

type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdatePasswordForm struct {
	Password string `json:"password"`
}

type FormatForm interface {
	validation.Validatable
	Format()
}

// Page 第几页
// Limit 一页最多
// Flag base 或者 custom之一
// Type get/post/header/cookie/ua/black_ip/white_ip 大小写无关
type RulePage struct {
	Page     uint   `json:"page"`
	Limit    uint   `json:"limit"`
	Flag     string `json:"flag"`
	Type     string `json:"type"`
	ServerID uint   `json:"server_id"`
	URIID    uint   `json:"uri_id"`
}

func (r *RulePage) Format() {
	r.Flag = strings.ToUpper(r.Flag)
	r.Type = strings.ToUpper(r.Type)
}

func (r RulePage) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ServerID, validation.Required.Error("server_id是必填的")),
		validation.Field(&r.Flag, validation.In("base", "BASE", "custom", "CUSTOM").Error("类型必须是base或者custom之一")),
		validation.Field(&r.Type, validation.In("GET", "POST", "BLACKIP", "COOKIE", "UA", "HEADER", "WHITEIP").Error("规则类型不合法")),
	)
}

// 必要字段 "Content" "ServerID" "Flag" "Type"
// 可选 "URIID"
type DeleteRuleForm struct {
	Rule
	ServerID uint `json:"server_id"`
	URIID    uint `json:"uri_id"`
}

type GetURIForm struct {
	ServerID uint `json:"server_id"`
}

func (g GetURIForm) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ServerID, validation.Required.Error("server_id是必填的")),
	)
}

type AddRuleForm struct {
	ServerID uint   `json:"server_id"`
	URIID    uint   `json:"uri_id"`
	Rules    []Rule `json:"rules"`
}

func (a AddRuleForm) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ServerID, validation.Required.Error("必须指定服务器ID")),
		validation.Field(&a.Rules),
	)
}

func (a *AddRuleForm) Format() {
	for index, rule := range a.Rules {
		rule.Format()
		if a.URIID != 0 {
			rule.Flag = "CUSTOM"
		} else {
			rule.Flag = "BASE"
		}
		a.Rules[index] = rule
	}
}

type GetServerForm struct {
	ServerID uint `json:"server_id"`
}

func (g GetServerForm) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ServerID, validation.Required.Error("server_id 是必填的")),
	)
}

type SwitchOperation struct {
	ServerID     uint   `json:"server_id"`
	URIID        uint   `json:"uri_id"`
	ConfigName   string `json:"config_name"`
	ConfigStatus bool   `json:"config_value"`
}

func (s SwitchOperation) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ServerID, validation.Required.Error("必须声明服务器ID")),
		validation.Field(&s.ConfigName, validation.By(RuleLimit)),
	)
}

func RuleLimit(item interface{}) error {
	configName := item.(string)
	allowed := ""
	for _, val := range AbbrFuncMap {
		allowed += "," + val
	}
	if !strings.Contains(allowed, configName) {
		return errors.New("未定义规则")
	} else {
		return nil
	}
}

func (s *SwitchOperation) Format() {
	if val, ok := AbbrFuncMap[s.ConfigName]; ok {
		s.ConfigName = val
	}
	return
}

type GetSwitchForm struct {
	ServerID uint `json:"server_id"`
	URIID    uint `json:"uri_id"`
}

func (g GetSwitchForm) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ServerID, validation.Required.Error("必须声明服务器ID")),
	)
}

type GetWafStatusForm struct {
	ServerID uint `json:"server_id" form:"server_id"`
}

func (g GetWafStatusForm) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.ServerID, validation.Required.Error("必须声明服务器ID")),
	)
}

type RuleListForm struct {
	ServerID uint   `json:"server_id"`
	URIID    uint   `json:"uri_id"`
	Content  string `json:"content"`
	Type     string `json:"type"`
}

func (r *RuleListForm) Format() {
	r.Type = strings.TrimSpace(strings.ToUpper(r.Type))
}

func (r RuleListForm) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ServerID, validation.Required.Error("必须声明服务器ID")),
		validation.Field(&r.Type, validation.Required.Error("规则类型是必须声明的"), validation.In("GET", "POST", "COOKIE", "HEADER", "UA", "BLACKIP", "WHITEIP").Error("必须是GET POST COOKIE HEADER UA BLACKIP WHITEIP 之一")),
	)
}
