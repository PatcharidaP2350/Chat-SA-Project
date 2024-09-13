import { useEffect, useState } from "react";
import { Avatar, Col, Form, Row, Space } from "antd";
import { UserOutlined } from '@ant-design/icons';
import { MemberInterface} from "../interfaces/IMember";
import {MessageInterface } from "../interfaces/IMessage";
import "./ChatSeller.css";
import { GetMember, SetMessage } from "../services/https";
import { SendOutlined } from '@ant-design/icons';

// function ChatSeller() {
//   const [Member, setMember] = useState<MemberInterface[]>([]);

//   const getMember = async () => {
//     let res = await GetMember(1);
//     if (res) {
//       setMember([res]); // สมมติว่าข้อมูลที่ได้จาก GetMember เป็น object ไม่ใช่ array
//     }
//   };

//   useEffect(() => {
//     getMember();
//   }, []); // เพิ่ม [] เพื่อป้องกันการเรียกใช้ซ้ำ

function ChatSeller() {
  const [Member, setMember] = useState<MemberInterface[]>([]);
  const [loading, setLoading] = useState(true); // สถานะสำหรับการโหลด

  const onFinish = async (values: MessageInterface) => {
    
    
    let res = await SetMessage(values);
    console.log(res);
    if (res) {
      messageApi.open({
        type: "success",
        content: "บันทึกข้อมูลสำเร็จ",
      });
      // setTimeout(function () {
      //   // navigate("/SellerHome");
      // }, 2000);
    } else {
      messageApi.open({
        type: "error",
        content: "เกิดข้อผิดพลาด !",
      });
    }
  };


    const getMember = async () => {
      setLoading(true); // เริ่มต้นการโหลด
      let res = await GetMember(1);
      if (res) {
        setMember([res]); // สมมติว่าข้อมูลที่ได้จาก GetMember เป็น object ไม่ใช่ array
      }
      setLoading(false); // โหลดเสร็จสิ้น
    };

    useEffect(() => {
      getMember();
    }, []);

    return (
      <>
      <Form 
        onFinish={onFinish}>
        <Row>
          <Col
            style={{
              borderRadius: "12px",
              marginLeft: "300px",
              padding: "24px",
              background: "#e2dfdf",
              height: "725px",
              width: "900px",
              position: "relative", // เพื่อให้สามารถใช้ position: absolute ด้านในได้
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
                display: "flex", // ใช้ Flexbox เพื่อจัดตำแหน่ง
                alignItems: "center", // จัดแนวข้อความและ Avatar
                color: "#ffffff", // สีของข้อความ
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
                    backgroundColor: "#ffff", // สีพื้นหลัง
                    borderColor: "#3b3b3b", // เส้นขอบ
                    color: "#3b3b3b", // สีข้อความ
                  }}
                />
              </Space>
              <div
                style={{
                  marginLeft: "16px",
                  color: "#ffffff", // เปลี่ยนสีข้อความเป็นสีขาว
                  marginTop: "-15px", // ปรับตำแหน่งข้อความให้สูงขึ้น
                }}
              >
                {/* <span>ชื่อ Member (ผู้ซื้อ)</span> ข้อความที่ต้องการใส่ */}
                <span>
                  {/* {Member.length > 0 ? `${Member[0].FirstName} ${Member[0].LastName}` : "ชื่อ Member (ผู้ซื้อ)"} */}
                  {loading
                    ? "กำลังโหลด..." // ข้อความในขณะโหลด
                    : Member.length > 0
                    ? `${Member[1].FirstName} ${Member[1].LastName}` // แสดงชื่อหลังจากโหลดสำเร็จ
                    : "ไม่พบข้อมูลผู้ใช้"} {/* แสดงข้อความเมื่อข้อมูลยังไม่มี */}
                </span> แสดงชื่อผู้ซื้อ
              </div>
            </Row>

            <div className="input" 
            style={{
              display: "flex", // ใช้ flexbox เพื่อจัดวาง input และ icon
              alignItems: "center", // จัดให้อยู่กึ่งกลางแนวตั้ง
              marginTop: "610px",
              width: "100%", // ให้ div ขยายเต็มพื้นที่
              position: "relative",
            }}
              >
              <input
                type="text"
                placeholder="พิมพ์ข้อความที่นี่"
                style={{
                  flex: 1, // ให้ input ขยายเต็มที่
                  padding: "10px",
                  borderRadius: "8px",
                  border: "1px solid #ccc",
                  marginRight: "8px", // เว้นระยะทางขวาเล็กน้อยเพื่อเว้นระยะไอคอน
                }}
              />
                <SendOutlined
                  style={{
                    cursor: "pointer",
                    fontSize: "24px", // ขนาดของไอคอน
                    color: "#000000", // สีของไอคอน
                  }}
                  onClick={() => {
                    // เพิ่มฟังก์ชันเมื่อคลิกที่ไอคอนนี้
                    console.log("ส่งข้อความ");
                  }}
                />
            </div>

            {/* แสดงชื่อ Member */}
            <div>
              {Member.map((item) => (
                <div key={item.ID}>
                  {item.FirstName} {item.LastName}
                </div>
              ))}
            </div>
          </Col>
        </Row>
      </Form>
      </>
  );
}

export default ChatSeller;
