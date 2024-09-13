import { useEffect, useState } from "react";
import { Avatar, Card, Col, Form, Row, Space, message } from "antd";
import { UserOutlined, SendOutlined, RedoOutlined  } from '@ant-design/icons';
import { MessageInterface } from "../interfaces/IMessage";
import "./ChatSeller.css";
import back from "../assets/back.png";

import { GetMember, SetMessage } from "../services/https";

interface Member {
  MemberID:   number
	Username:    string
	Password:    string
	Email:       string
	FirstName:  string
	LastName:    string
	PhoneNumber: string
	Address:     string
	PicProfile:  string
}


function ChatSeller() {
  const [Member, setMember] = useState<Member[]>([]);
  const [loading, setLoading] = useState(true);
  const [messageApi, contextHolder] = message.useMessage();

  const onFinish = async (values: MessageInterface) => {
    let res = await SetMessage(values);
    console.log(res);
    if (res) {
      messageApi.open({
        type: "success",
        content: "บันทึกข้อมูลสำเร็จ",
      });
    } else {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาด !",
      });
    }
  };

  const getMember = async () => {
    setLoading(true);
    let data = await GetMember(1); // เรียก API เพื่อดึงข้อมูลของสมาชิก
    console.log(data);
    if (Array.isArray(data)) {
      setMember(data); // ใช้ array ตรงๆ ถ้า data เป็น array
    } else {
      setMember([data]); // ถ้า data เป็น object ให้ใช้ [data] เพื่อทำให้เป็น array
    }
    setLoading(false);
  };
  

  useEffect(() => {
    getMember();
  }, []);

  return (
    <>
      {contextHolder}
      <Form onFinish={onFinish}>
        <div className="iconHomeChatSeller">
          <img
                  src={back}
                  alt="backarrow"
                  // onClick={handleBacktoHome}
                  style={{
                    width: "35px",
                    height:"35px",
                    cursor: "pointer",
                    marginLeft: "1400px",
                  }}
            />
             <Avatar
                    size={64}
                    icon={<RedoOutlined />}
                    style={{
                      marginTop: "-5px",
                      width: "70px",
                      height: "45px",
                      backgroundColor: "#0000",
                      color: "#000",
                    }}
                  />
        </div>
          <Row>
            <Col
              style={{
                borderRadius: "12px",
                marginLeft: "300px",
                padding: "24px",
                background: "#e2dfdf",
                height: "725px",
                width: "900px",
                position: "relative",
                marginTop:"-42px"
              }}
            >
              <Row
                style={{
                  borderRadius: "12px 12px 0 0",
                  marginLeft: "-24px",
                  marginTop: "-23.7px",
                  padding: "24px",
                  background: "#3b3b3b",
                  height: "60px",
                  width: "900px",
                  display: "flex",
                  alignItems: "center",
                  color: "#ffffff",
                }}
              >
                <Space wrap size={16}>
                  <Avatar
                    size={64}
                    icon={<UserOutlined />}
                    style={{
                      marginTop: "-16px",
                      width: "45px",
                      height: "45px",
                      backgroundColor: "#ffff",
                      borderColor: "#3b3b3b",
                      color: "#3b3b3b",
                    }}
                  />
                </Space>
                <div
                  style={{
                    marginLeft: "16px",
                    color: "#ffffff",
                    marginTop: "-15px",
                  }}
                >
                  <span>
                    {loading
                      ? "กำลังโหลด..."
                      : Member.length > 0
                      ? `${Member[0].FirstName} ${Member[0].LastName}`
                      : "ไม่พบข้อมูลผู้ใช้"}
                  </span>
                  
                </div>              

              </Row>

              {/* ส่วนสำหรับ input และไอคอน */}
              <div className="input"
                style={{
                  display: "flex",
                  alignItems: "center",
                  marginTop: "610px",
                  width: "100%",
                  position: "relative",
                }}
              >
                <input
                  type="text"
                  placeholder="พิมพ์ข้อความที่นี่"
                  style={{
                    flex: 1,
                    padding: "10px",
                    borderRadius: "8px",
                    border: "1px solid #ccc",
                    marginRight: "8px",
                  }}
                />
                <SendOutlined
                  style={{
                    cursor: "pointer",
                    fontSize: "24px",
                    color: "#000000",
                  }}
                  onClick={() => {
                    console.log("ส่งข้อความ");
                  }}
                />
              </div>

              {/* แสดงชื่อ Member */}
              
            </Col>
          </Row>
      </Form>
    </>
  );
}

export default ChatSeller;