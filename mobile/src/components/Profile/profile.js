import React, { Component } from 'react';

import { 
    ContainerProfile,
    ImgProfileConfig, 
    TextNameUser, 
    ViewConfig, 
    TextConfig,
    ViewConfigList,
    TextConfigList,
    ButtonConfigList,
    ButtonOut,
    ViewButtonOut,
    TextBoxButtonOut,
    ViewRodape,
    TextBoxRodape
} from "./styles_profile";
import { ImageBackground, StyleSheet, Text} from 'react-native';

export default class Profile extends Component{  

    render(){  
        return( 
       <ContainerProfile>
           <ImageBackground source={require('../../assets/fundo.png')} style={styles.backgroundImage}>

                <ImgProfileConfig source={require('../../assets/perfil_image.png')}/>
                <TextNameUser>Offar</TextNameUser>
    
                <ViewConfig>
                    <TextConfig>Configurações</TextConfig>
                </ViewConfig>

                <ViewConfigList>
                    <ButtonConfigList><TextConfigList>Sua Ficha</TextConfigList></ButtonConfigList>
                    <ButtonConfigList><TextConfigList>Editar Perfil</TextConfigList></ButtonConfigList>
                    <ButtonConfigList><TextConfigList>Configurações de E-mail</TextConfigList></ButtonConfigList>
                    <ButtonConfigList><TextConfigList>Manual do Jogador</TextConfigList></ButtonConfigList>
                </ViewConfigList>

                <ButtonOut  onPress={() => this.props.navigation.navigate('Login') } >
                    <ViewButtonOut>
                        <TextBoxButtonOut>Sair</TextBoxButtonOut>
                    </ViewButtonOut>
                </ButtonOut>
                <ViewRodape>
                    <TextBoxRodape>Feito com ♡ por Cubisme Design Team</TextBoxRodape>
                    <TextBoxRodape>Version: 0.3</TextBoxRodape>
                </ViewRodape>

           </ImageBackground>                                 
       </ContainerProfile>    
  
  )  
}  
} 

const styles = StyleSheet.create({
    backgroundImage: {
        height: '100%',
        width: '100%',
        alignItems: 'center',
        resizeMode: 'cover',
  }
  });

Profile.navigationOptions = {
    header: null
  }