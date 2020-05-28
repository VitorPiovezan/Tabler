import styled from 'styled-components/native';
/* Styles Profile */
export const ContainerProfile = styled.View`
    flex: auto;
    align-items: center;
    background-color: #D9BA8E;
`;
export const ImgProfileConfig = styled.Image`
    width: 150px;
    height: 150px;
    border-radius: 100;
    margin: 50px 0px 0px 0px;
`;
export const TextNameUser = styled.Text`
    color: #5e3200;
    font-weight: bold;
    font-size: 26;
    padding: 20px
`;
export const ViewConfig = styled.View`
    border-style: solid;
    border-top-color: #5e3200;
    border-top-width: 1px;
    height: 60px;
    width: 40%;
    align-items: center;
    justify-content: center;
`;
export const TextConfig = styled.Text`
    color: #5e3200;
    font-weight: bold;
    font-size: 20;
`;
export const ViewConfigList = styled.View`
    border-style: solid;
    border-top-color: #5e3200;
    border-top-width: 2px;
    height: 100%;
    width: 80%;
`;
export const ButtonConfigList = styled.TouchableOpacity`
    border-style: solid;
    border-bottom-width: 1px;
    border-bottom-color: #5e3200;
    padding: 15px 0 0 15px;
    justify-content: center;
`;
export const TextConfigList = styled.Text`
    color: #5E3200;
    font-size: 18;
    padding-bottom: 15px;
`;

/* Rodapé */

export const ViewRodape = styled.View`
    border-style: solid;
    border-top-width: 1px;
    border-top-color: #5E3200
    width: 100%;
    height: 70px;
    bottom: 0;
    position: absolute;
    align-items: center;
    justify-content: center;
`;
export const TextBoxRodape = styled.Text`
    font-size: 12;
    color: #5E3200;
    margin-bottom: 4px;
`;
export const ViewButtonOut = styled.View`
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 50px;
    border-style: solid;
    border-width: 1px;
    border-color: #5E3200;
    border-radius: 30px;
`;
export const ButtonOut = styled.TouchableOpacity`
    justify-content: center;
    align-items: center;
    width: 50%;
    height: 40px;
    position: absolute;
    bottom: 100;
`;
export const TextBoxButtonOut = styled.Text`
    font-size: 20;
    color: #5E3200;
`;