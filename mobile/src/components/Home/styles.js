import styled from 'styled-components/native';


export const Container = styled.View`
    flex: auto;
    justify-content: center;
    align-items: center;
    background-color: #D9BA8E;
`
export const ContainerHome = styled.View`
    justify-content: center;
    align-items: center;
`

export const Card = styled.View`
    width: 90%;
    height:90%;
    background-color: #303030;
    border-radius: 15px;
`
export const Image = styled.Image`
    width: 100%;
    height: 65%;
    border-top-left-radius: 7px;
    border-top-right-radius: 7px;
`
export const Nome = styled.Text`
    font-weight: bold;
    color: #fff;
    font-size: 24;
    padding: 3px 0px 5px 10px;

`

export const ContainerIcons = styled.View`
    justify-content: center;
    align-items: center;
    position: absolute;
    bottom: 0;
`

export const ButtonStyles = styled.View`
    flex-direction: row;
    justify-content: space-between;
    width: 100%;
    padding: 5%;
`
export const Button = styled.TouchableOpacity`
    background-color: ${props => props.color};
    align-items: center;
    justify-content: center;
    width: 50%;
    height: 42px;
    border-radius: 7px;
`
export const IconMath = styled.Image`

`
export const InfoProfile = styled.Text`
    color: #fff;
    font-size: 16;
    padding: 3px 0px 5px 10px;
`
export const Info = styled.View`

`
export const TextExemple = styled.Text`
    font-weight: bold;
    font-size: 24;
    color: white;
`

export const InputChat = styled.TextInput`
    height: 50px;
    width: 95%;
    border-radius: 7px;
    background-color: #fff;
    padding: 15px;
    font-size: 16;
    position: absolute;
    bottom: 10;
    
`

export const ViewChat = styled.View`
    background-color: #303030;
    border-radius: 7px;
    width: 95%;
    max-height: 79%;
    min-height: 50%;
    margin-bottom: 10px;
    padding: 10px 0 10px 0;
`

export const MsgBoxUser = styled.View`
    margin: 0px 10px 10px 10px;
    background-color: #aaa;
    height: 50px;
    flex-direction: row;
    align-items: center;
    border-radius: 7;
    padding: 10px;
`

export const MsgBoxProfile = styled.View`
    margin: 0px 10px 10px 10px;
    background-color: #505050;
    height: 50px;
    flex-direction: row;
    align-items: center;
    border-radius: 7;
    padding: 10px;
    justify-content: flex-end;
`

export const Img = styled.Image`
    width: 40px;
    height: 40px;
    border-radius: 100;
`
export const ImgProfile = styled.Image`
    width: 40px;
    height: 40px;
    border-radius: 100;
`

export const TextBox = styled.Text`
    padding: 10px;
    color: #000;
`

export const TextBoxProfile = styled.Text`
    padding: 10px;
    color: #fff;
`

export const ContainerChat = styled.View`
    align-items: center;
    background-color: #D9BA8E;
    flex:auto;
`

export const ContainerFeed = styled.View`
    align-items: flex-start;
    background-color: #D9BA8E;
    flex: auto;
`

export const ViewMatch = styled.View`
    width: 95%;
    height: 75px;
    flex-direction: row;
    padding: 5px;
    background-color: #303030;
    border-radius: 7;
    margin: 10px 10px 10px 10px;
`
export const UserMatch = styled.View`
    flex-direction: column;
    padding: 10px;
    justify-content: center;
    align-items: center;
    
`

export const UserMatchTouch = styled.TouchableOpacity`
    flex-direction: column;
    padding: 10px;
    justify-content: center;
    align-items: center;
    
`

export const ImgUser = styled.Image`
    width: 45px;
    height: 45px;
    border-radius: 100;
`
export const NameUser = styled.Text`
    color: #fff;
    font-weight: bold;
    font-size: 12;
`

export const UserMathButton = styled.TouchableOpacity`
    margin-top: -7px;
`

export const ViewGames = styled.View`
    width: 95%;
    height: 75px;
    flex-direction: row;
    padding: 5px;
    background-color: #303030;
    border-radius: 7;
    margin: 10px;
`
export const TextGames = styled.Text`
    font-weight: bold;
    font-size: 20;
    color: #fff;
    padding: 20px 0px 0px 12px;
`
export const TextMessages = styled.Text`
    font-weight: bold;
    font-size: 20;
    color: #fff;
    padding: 0px 0px 10px 12px;
`
export const ViewChatsUsers = styled.View`
    
    width: 100%;
    height: 100%;
    align-items: flex-start;
`
export const ChatUsers = styled.View`
    border-bottom-width: 1;
    border-color: #303030;
    width: 100%;
    height: 90px;   
    flex-direction: row;
    align-items: center;
    padding: 10px;
`
export const ImgChatUser = styled.Image`
    width: 65px;
    height: 65px;
    border-radius: 100;
`

export const TextBoxChatUser = styled.View`
    justify-content: center;
`

export const TextBoxChatUserName = styled.Text`
    padding: 10px 0 0 10px;
    color: #fff;
    font-weight: bold;
    font-size: 20;
`
export const TextBoxChatUserLastMassage = styled.Text`
    padding: 3px 0 10px 10px;
    color: #fff;
    font-size: 16;
`
export const ButtonChatUser = styled.TouchableOpacity`
    width: 100%;
`




//Styles Home

export const ContainerProfile = styled.View`
    flex: auto;
    align-items: center;
    background-color: #D9BA8E;
`;
export const ViewHeaderHome = styled.View`
    border-bottom-left-radius: 100;
    border-bottom-right-radius: 100;
    width: 100%;
    min-height: 200px;
    background-color: #5e3200;

`;
export const IconsProfile = styled.Image`
    width: 30px;
    height: 30px;
`;
export const ButtonProfile = styled.TouchableOpacity`
    justify-content: center;
    align-items: center;
    border-radius: 100;
    width: 75px;
    height: 75px;
    background-color: #DDBC91;
`;
export const ViewConfigsProfile = styled.View`
    justify-content: space-between;
    align-items: center;
    margin: auto;
    width: 70%;
    height: 160px;
    flex-direction: row;
`;
export const TextBoxNameProfile = styled.Text`
    color: #DDBC91;
    font-size: 20;
`;
export const ViewTextProfile = styled.View`
    width: 100%;
    align-items: center;
    height: 40px;    
`;
export const ViewContentProfile = styled.View`
    width: 100%;
    background-color: #303030;
    height: 140px;
    margin-top: 15px;
    justify-content: center;
`;
export const TextBoxContentProfile = styled.Text`
    font-size: 20;
    color: #fff;
    font-weight: bold;
    margin-left: 15px;
    margin-top: 10px;
`;
export const ViewGamesProfile = styled.View`
    width: 100%;
    height: 90px;
    flex-direction: row;
    padding: 5px;
    border-radius: 7;
    margin: 10px 0 10px 0;
`;
export const ImgUserProfile = styled.Image`
    width: 60px;
    height: 60px;
    border-radius: 100;
`;
export const NameUserProfile = styled.Text`
    color: #fff;
    font-weight: bold;
    font-size: 14;
    margin-top: 5px;
`;
export const ViewRodape = styled.View`
    width: 100%;
    height: 70px;
    bottom: 0;
    background-color: #303030;
    position: absolute;
    align-items: center;
    justify-content: center;
`;
export const TextBoxRodape = styled.Text`
    font-size: 12;
    color: #ccc;
    margin-bottom: 4px;
`;
export const ViewButtonOut = styled.View`
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 60px;
    background-color: #303030;
    border-top-left-radius: 30px;
    border-top-right-radius: 30px;
`;
export const ButtonOut = styled.TouchableOpacity`

    justify-content: center;
    align-items: center;
    width: 100%;
    height: 60px;
    position: absolute;
    bottom: 75;
`;
export const TextBoxButtonOut = styled.Text`
    font-size: 20;
    color: #ddd;
`;

export const CardModal = styled.TouchableOpacity`
    width: 80%;
    height: 60%;
    background-color: #202020;
    border-radius: 7px;
    border: 1px;
    border-color: rgba(10, 10, 10, 0.75);
`;

export const ImageModal = styled.Image`
    width: 100%;
    height: 30%;
    border-top-left-radius: 7px;
    border-top-right-radius: 7px;
`;

export const NomeModal = styled.Text`
    font-weight: bold;
    color: #fff;
    font-size: 24;
    padding: 10px 15px 5px 15px;

`;

export const InfoModal = styled.Text`
    color: #fff;
    font-size: 16;
    text-align: justify;
    padding: 3px 15px 5px 15px;
`;

export const ButtonModal = styled.TouchableOpacity`
    background-color: #ac58aa;
    width: 50%;
    height: 40px;
    align-items: center;
    flex-direction: column;
    border-radius: 4px;
    position: absolute;
	bottom: 0px;
`;

export const TextModal = styled.Text`
    padding-top: 8px;
    font-size: 18px;
    color: #fff;
    font-weight: bold;
`;

export const InfoModalView = styled.View`
    align-items: center;
    position: relative;
    height: 60%;
`
