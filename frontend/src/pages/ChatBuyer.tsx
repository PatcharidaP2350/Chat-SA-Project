import {Col, Row} from "antd"; 
// import Column from "antd/es/table/Column";
import { UserOutlined } from '@ant-design/icons';
import { Avatar, Space } from 'antd';
// import React from 'react';

function ChatBuyer (){

return (
        <>
        
    <Row>
          <Col         
            style={{
              borderRadius: "12px",
              marginLeft: "600px",
              padding: "24px",
              background: "#e2dfdf",
              height: "725px",
              width:"900px",
              
            }}>
                <Row style={{
                        borderRadius: "12px 12px 0 0",
                        marginLeft: "-24px",
                        marginTop: "-23.7px",
                        padding: "24px",
                        background: "#3b3b3b",
                        height: "60px",
                        width:"900px",
                        
                      }}>

                    <Space wrap size={16 }>
                    <Avatar size={64} icon={<UserOutlined />} 
                        style={{
                            marginTop: "-20px",
                            width: "50px",
                            height: "50px",
                            backgroundColor: "#ffff",           //สีพื้นหลัง
                            borderColor: "#3b3b3b",          //เส้นขอบ
                            color: "#3b3b3b",      //สีข้อความ
                        }}/>

                    </Space>
                </Row>
          <div className="input" style={{ marginTop: "20px" }}>
            <input
              type="text"
              placeholder="พิมพ์ข้อความที่นี่"
              style={{
                width: "100%",
                padding: "10px",
                borderRadius: "8px",
                border: "1px solid #ccc",
                marginTop: "590px",
                marginLeft: "-11px",
              }}
            />
          </div>
          </Col>
        </Row>
        </>
      );
    };




export default ChatBuyer