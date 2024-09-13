import { RoomChatInterface } from "../../interfaces/IRoomChat";
import { MessageInterface } from "../../interfaces/IMessage";
// import { MemberInterface } from "../../interfaces/IMember;

const apiUrl = "http://localhost:8080";

async function CreateRoomChat(data: RoomChatInterface) {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };
  
    let res = await fetch(`${apiUrl}/roomchat`, requestOptions)
      .then((res) => {
        if (res.status == 201) {
          return res.json();
        } else {
          return false;
        }
      });
  
    return res;
  }



  async function SetMessage(data: MessageInterface) {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    };
  
    let res = await fetch(`${apiUrl}/message`, requestOptions)
      .then((res) => {
        if (res.status == 201) {
          return res.json();
        } else {
          return false;
        }
      });
  
    return res;
  }



  async function GetMessage(id:Number | undefined) {
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
  
    let res = await fetch(`${apiUrl}/message/${id}`, requestOptions)
      .then((res) => {
        if (res.status == 200) {
          return res.json();
        } else {
          return false;
        }
      });
  
    return res;
  }  

  async function GetRoomChat() {
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
  
    let res = await fetch(`${apiUrl}/roomchat`, requestOptions)
      .then((res) => {
        if (res.status == 200) {
          return res.json();
        } else {
          return false;
        }
      });
  
    return res;
  }  
  async function GetMember(member_id:Number | undefined) {
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
  
    let res = await fetch(`${apiUrl}/member/${member_id}`, requestOptions)
      .then((res) => {
        if (res.status == 200) {
          return res.json();
        } else {
          return false;
        }
      });
  
    return res;
  }  

  async function GetSeller(seller_id:Number | undefined) {
    const requestOptions = {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    };
  
    let res = await fetch(`${apiUrl}/seller/${seller_id}`, requestOptions)
      .then((res) => {
        if (res.status == 200) {
          return res.json();
        } else {
          return false;
        }
      });
  
    return res;
  }  


  export {
    CreateRoomChat,
    SetMessage, 
    GetMessage,
    GetRoomChat,
    GetMember,
    GetSeller,
  };

