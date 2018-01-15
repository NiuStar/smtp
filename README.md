# smtp

#使用方法：

    func SendMail(to,subject,body string) bool {
    
        user := "a24802117@hotmail.com"
        password := "***********"
        host := "smtp.live.com:25"
    
        mail := smtp.NewMailAuth(user,password,host)
    
        //to := "24802117@qq.com"
    
        //subject := "又有新消息了"
    
        //body := `订单发生变更了,用户***提交了一个新的订单,需求为四川省成都市武侯区的五星级酒店豪华标准间一间,于2016年7月22日入住,请尽快处理`
    
        err := mail.SendMail(to,subject,body,"")
        if err != nil {
            fmt.Println("Send mail error!")
            fmt.Println(err)
            return false
        } else {
            fmt.Println("Send mail success!")
            return true
        }
    }