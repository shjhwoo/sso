import React from "react";
import { useState } from "react";
import {
  Button,
  Input,
  LoginSection,
  SaveInput,
  SubSection,
} from "./components/login";
import axios from "axios";

function App() {
  const [loginInfo, setLoginInfo] = useState({
    hospital_code: "",
    userid: "",
    password: "",
  });

  const handleInputChange =
    (key: string) => (e: React.ChangeEvent<HTMLInputElement>) => {
      setLoginInfo({ ...loginInfo, [key]: e.target.value });
    };

  const handleLoginRequest = () => {
    console.log("???");
    axios.post("http://localhost:8080/login", loginInfo, {
      withCredentials: true,
    });
  };

  return (
    <div className="App">
      <h1>통합회원 로그인</h1>
      <LoginSection>
        <Input>
          <input
            placeholder="기관코드"
            onChange={handleInputChange("hospital_code")}
            value={loginInfo.hospital_code}
          ></input>
        </Input>
        <Input>
          <input
            placeholder="아이디"
            onChange={handleInputChange("userid")}
            value={loginInfo.userid}
          ></input>
        </Input>
        <Input>
          <input
            placeholder="비밀번호"
            type="password"
            onChange={handleInputChange("password")}
            onKeyUp={(e) => (e.key === "Enter" ? handleLoginRequest() : null)}
            value={loginInfo.password}
          ></input>
        </Input>
        <Button onClick={handleLoginRequest}>통합회원 로그인</Button>
        <SaveInput>
          <input className="save-orgcode" type="checkbox"></input>
          <label htmlFor="save-orgcode">기관코드 저장</label>
          <input className="save-id" type="checkbox"></input>
          <label htmlFor="save-id">아이디 저장</label>
          <input className="save-password" type="checkbox"></input>
          <label htmlFor="save-password">비밀번호 저장</label>
        </SaveInput>
      </LoginSection>
      <SubSection>
        <p>아이디/비밀번호 찾기</p>
        <p>통합회원가입</p>
      </SubSection>
    </div>
  );
}

export default App;
