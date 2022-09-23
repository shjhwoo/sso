import { createGlobalStyle } from "styled-components";

const GlobalStyle = createGlobalStyle`
html,body{
  //border: 1px solid purple;
  box-sizing: border-box;
  margin: 0;
  width: 100%;
  height: 100%;
}
*, *:before, *:after {
  box-sizing: inherit;
}
#root {
  width: 100%;
  height: 100%;
  background-color: #0E9BC1;
  background-image: url(images/gopher.png)
  //border: 3px solid green;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.App{
  //border: 1px solid orange;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}
`;

export default GlobalStyle;
