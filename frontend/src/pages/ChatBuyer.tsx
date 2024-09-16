import { SendOutlined, UserOutlined, RedoOutlined } from '@ant-design/icons';
import { Avatar, Col, Form, Row, Space, message } from "antd";
import { useEffect, useRef, useState } from "react";
import { MessageInterface } from "../interfaces/IMessage";
import back from "../assets/back.png";
import { CreateMessage, GetMember, GetMemberBySeller, GetMessage } from "../services/https";
import "./ChatSeller.css";

interface Member {
  MemberID: number;
  Username: string;
  Password: string;
  Email: string;
  FirstName: string;
  LastName: string;
  PhoneNumber: string;
  Address: string;
  ProfilePic: string;
}



interface Memberbyseller {
  MemberID: number;
  Username: string;
  Password: string;
  Email: string;
  FirstName: string;
  LastName: string;
  PhoneNumber: string;
  Address: string;
  ProfilePic: string;
}


function ChatBuyer() {
  const [member, setMember] = useState<Member | null>(null);
  const [memberbyseller, setMemberBySeller] = useState<Memberbyseller | null>(null);
  const [messages, setMessages] = useState<MessageInterface[]>([]);
  const [loading, setLoading] = useState(true);
  const [messageApi, contextHolder] = message.useMessage();
  const [inputMessage, setInputMessage] = useState(""); // State สำหรับข้อความที่กรอก
  const messagesEndRef = useRef<HTMLDivElement>(null); // ใช้เพื่อเลื่อนดูข้อความล่าสุด

  const roomChatID = 1; // ห้องแชทที่ใช้
  const senderID = 2; // ID ของผู้ส่งข้อความ
  const sellerID = 1; // ID ของผู้ส่งข้อความ

  const onFinish = async () => {
    if (!inputMessage.trim()) {
      messageApi.open({
        type: "error",
        content: "กรุณากรอกข้อความ",
      });
      return;
    }

    const messageData: MessageInterface = {
      room_chat_id: roomChatID,
      content: inputMessage,
      sender_id: senderID,
    };

    try {
      const res = await CreateMessage(messageData);
      console.log("Message creation response:", res); // ตรวจสอบการตอบสนอง
      if (res) {
        messageApi.open({
          type: "success",
          content: "บันทึกข้อความสำเร็จ",
        });
        setInputMessage("");
        fetchMessages(); // ดึงข้อความใหม่หลังจากส่งข้อความ
      } else {
        messageApi.open({
          type: "error",
          content: "เกิดข้อผิดพลาด !",
        });
      }
    } catch (error) {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาดในการส่งข้อความ",
      });
    }
  };

  const fetchMessages = async () => {
    setLoading(true);
    try {
      const data = await GetMessage(roomChatID); // เรียก API เพื่อดึงข้อความ
      setMessages(Array.isArray(data) ? data : [data]);
    } catch (error) {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาดในการดึงข้อความ",
      });
    }
    setLoading(false);
  };

  const getMember = async () => {
    setLoading(true);
    try {
      const data = await GetMember(senderID);
      if (data) {
        setMember(data);
      } else {
        messageApi.open({
          type: "error",
          content: "ไม่พบข้อมูลสมาชิก",
        });
      }
    } catch (error) {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาดในการดึงข้อมูลสมาชิก",
      });
    }
    setLoading(false);
  };

  const getMemberBySeller = async () => {
    setLoading(true);
    try {
      const dataMember = await GetMemberBySeller(sellerID); // เรียก API เพื่อดึงข้อมูลของ memberbyseller
      console.log("datamember", dataMember);
      setMemberBySeller(dataMember);
    } catch (error) {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาดในการดึงข้อมูล memberbyseller",
      });
    }
    setLoading(false);
  };

  useEffect(() => {
    getMember();
    getMemberBySeller();
    fetchMessages(); // ดึงข้อความทั้งหมดเมื่อ component ถูกโหลด
  }, []);

  useEffect(() => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" }); // เลื่อนดูข้อความล่าสุด
  }, [messages]);

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
              height: "35px",
              cursor: "pointer",
              marginLeft: "1400px",
            }}
          />
          <Avatar
            size={64}
            icon={<RedoOutlined />}
            style={{
              marginTop: "-75px",
              width: "70px",
              height: "45px",
              backgroundColor: "#0000",
              color: "#000",
              marginLeft: "1435px",
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
              marginTop: "-65px"
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
                  src={member?.ProfilePic}
                  icon={!member?.ProfilePic ? <UserOutlined /> : undefined}
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
                    : memberbyseller
                      ? `${memberbyseller.FirstName} ${memberbyseller.LastName}`
                      : "ไม่พบข้อมูลผู้ใช้"}
                </span>
              </div>
            </Row>

            {/* แสดงข้อความ */}
            <div className="messages" style={{
              marginTop: "16px",
              overflowY: "auto",
              flex: 1,
              display: "flex",
              flexDirection: "column", // ให้ข้อความใหม่แสดงที่ด้านล่าง
              height: "calc(100% - 80px)",
            }}>
              {messages.map((msg) => (
                <div key={msg.ID} style={{
                  marginBottom: "10px",
                  textAlign: msg.SenderID === senderID ? "right" : "left", // ตรวจสอบ ID ผู้ส่ง
                  padding: "10px",
                  backgroundColor: msg.SenderID === senderID ? "#d1ffd1" : "#f1f1f1", // สีพื้นหลังต่างกัน
                  borderRadius: "8px",
                  display: "inline-block",
                  maxWidth: "60%", // จำกัดความกว้างสูงสุด
                  wordBreak: "break-word", // ตัดคำหากข้อความยาวเกินไป
                  marginLeft: msg.SenderID === senderID ? "auto" : "0", // จัดข้อความไปทางขวาหรือซ้าย
                }}>
                  {msg.Content}
                </div>
              ))}
              <div ref={messagesEndRef} /> {/* ใช้เพื่อเลื่อนดูข้อความล่าสุด */}
            </div>

            {/* ส่วนสำหรับ input และไอคอน */}
            <div className="input"
              style={{
                display: "flex",
                alignItems: "center",
                marginTop: "auto",
                width: "100%",
                position: "relative",
              }}
            >
              <Form.Item
                name="message"
                style={{
                  width: "calc(100% - 80px)",
                  margin: "0",
                  padding: "0",
                }}
              >
                <input
                  type="text"
                  value={inputMessage}
                  onChange={(e) => setInputMessage(e.target.value)}
                  placeholder="พิมพ์ข้อความ..."
                  style={{
                    width: "100%",
                    padding: "8px",
                    borderRadius: "4px",
                    border: "1px solid #ddd",
                    boxSizing: "border-box",
                  }}
                />
              </Form.Item>
              <div
                onClick={onFinish}
                style={{
                  position: "absolute",
                  right: "0",
                  padding: "8px",
                  cursor: "pointer",
                }}
              >
                <SendOutlined />
              </div>
            </div>
          </Col>
        </Row>
      </Form>
    </>
  );
}

export default ChatBuyer;
