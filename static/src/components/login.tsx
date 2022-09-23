import styled from "styled-components";

export const LoginSection = styled.section`
  border: 1px solid pink;
  background-color: #dbdbdb;
  padding: 0.8rem;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 250px;
  width: 360px;
`;

export const Input = styled.section`
  //border: 1px solid red;
  margin-bottom: 1rem;
  width: 100%;
  height: 40px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  > label {
    font-size: 1rem;
    font-weight: bold;
  }
  > input {
    &:focus {
      outline: none;
    }
    border: 1px solid gray;
    padding-left: 0.5rem;
    width: 100%;
    height: 100%;
    background-color: #c5e5f6;
    font-size: 1rem;
  }
`;

export const SaveInput = styled.section`
  margin-top: 0.5rem;
  width: 100%;
  align-items: center;
  > label {
    margin-left: 0;
    margin-right: 0.5rem;
    font-size: 0.8rem;
    font-weight: bold;
    color: #3f3f3f;
  }
`;

export const Button = styled.button`
  border: none;
  background-color: #787878;
  color: white;
  font-size: 1rem;
  font-weight: 500;
  width: 100%;
  height: 40px;
  &:focus {
    outline: none;
  }
  &:hover {
    background-color: #cdcdcd;
    background-image: url(https://img1.daumcdn.net/thumb/R800x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FAaHEH%2FbtqDTktG4pV%2FZHyelcepRnBVrhPKqhHxpk%2Fimg.png);
    background-position: right;
    background-repeat: no-repeat;
    background-size: contain;
  }
`;

export const SubSection = styled.section`
  //border: 1px solid green;
  //padding: 0rem 0.5rem;
  width: 360px;
  display: flex;
  justify-content: space-between;
  color: #3f3f3f;
  font-size: 0.8rem;
  > p {
    font-weight: bold;
  }
`;
