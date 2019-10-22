package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type MrzdCashinLog struct {
	Id         int     `orm:"column(id);auto"`
	UserId     int     `orm:"column(user_id)" description:"用户ID"`
	ChannelId  int     `orm:"column(channel_id)" description:"渠道号"`
	RealName   string  `orm:"column(real_name);size(50);null" description:"提现人昵称"`
	Phone      string  `orm:"column(phone);size(20)" description:"提现人手机号"`
	Alipay     string  `orm:"column(alipay);size(50)" description:"支付宝账号 "`
	Money      float64 `orm:"column(money);digits(7);decimals(2)" description:"提现金额;"`
	Status     int8    `orm:"column(status)" description:"支付状态；1：未支付；2：已支付；3：拒绝支付"`
	Reason     string  `orm:"column(reason);size(255);null" description:"拒绝支付原因"`
	CreateDate string  `orm:"column(create_date);size(20)" description:"提现日期：2018-05-10"`
	CreateTime int     `orm:"column(create_time)" description:"提现时间；unix时间戳"`
	PayDate    string  `orm:"column(pay_date);size(20);null" description:"支付日期"`
	PayTime    int     `orm:"column(pay_time);null" description:"支付时间"`
}

func (t *MrzdCashinLog) TableName() string {
	return "mrzd_cashin_log"
}

func init() {
	orm.RegisterModel(new(MrzdCashinLog))
}

// AddMrzdCashinLog insert a new MrzdCashinLog into database and returns
// last inserted Id on success.
func AddMrzdCashinLog(m *MrzdCashinLog) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetMrzdCashinLogById retrieves MrzdCashinLog by Id. Returns error if
// Id doesn't exist
func GetMrzdCashinLogById(id int) (v *MrzdCashinLog, err error) {
	o := orm.NewOrm()
	v = &MrzdCashinLog{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMrzdCashinLog retrieves all MrzdCashinLog matches certain condition. Returns empty list if
// no records exist
func GetAllMrzdCashinLog(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(MrzdCashinLog))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []MrzdCashinLog
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMrzdCashinLog updates MrzdCashinLog by Id and returns error if
// the record to be updated doesn't exist
func UpdateMrzdCashinLogById(m *MrzdCashinLog) (err error) {
	o := orm.NewOrm()
	v := MrzdCashinLog{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMrzdCashinLog deletes MrzdCashinLog by Id and returns error if
// the record to be deleted doesn't exist
func DeleteMrzdCashinLog(id int) (err error) {
	o := orm.NewOrm()
	v := MrzdCashinLog{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&MrzdCashinLog{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
