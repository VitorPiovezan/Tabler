import React, { Component } from 'react';

import {
    ContainerHome,
    ViewHeaderHome,
    ViewConfigsProfile,
    ButtonProfile,
    IconsProfile,
    ViewTextProfile,
    ImgHomeConfig,
    TextBoxNameProfile,
    TitleHome,
    Input,
    ViewRoom,
    ViewOpenRoom,
    ViewTitles,
    TitleRoom,
    PlayersRoom,
    ButtonRoom,
    TextButtonRoom,
    ViewButtonRoom
} from './styles';

import { ScrollView } from 'react-native-gesture-handler';
import { StyleSheet, ImageBackground } from 'react-native';

export default class Home extends Component{  
  render(){  
      return(   
      <ContainerHome>
        <ImageBackground source={require('../../assets/fundo.png')} style={styles.backgroundImage} >

        <ViewHeaderHome>
                <ViewConfigsProfile>
                    <ButtonProfile onPress={() => this.props.navigation.navigate('Config') }>
                        <IconsProfile source={require('../../assets/configIcon.png')}/>
                    </ButtonProfile>
                    <ImgHomeConfig source={require('../../assets/icon.png')}/>
                    <ButtonProfile onPress={() => this.props.navigation.navigate('EditProfile') } >
                        <IconsProfile source={require('../../assets/perfilIcon.png')}/>
                    </ButtonProfile>
                </ViewConfigsProfile>

                <ViewTextProfile>
                    <TextBoxNameProfile>Bem Vindo Offar</TextBoxNameProfile>
                </ViewTextProfile>
            </ViewHeaderHome>

        <TitleHome>Salas Abertas</TitleHome>
        <Input placeholder="Pesquisar salas.." />

        <ViewOpenRoom>
            <ScrollView>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Descent to Avernus</TitleRoom>
                        <PlayersRoom>6/10 Jogadores | 1/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 1</TitleRoom>
                        <PlayersRoom>4/10 Jogadores | 1/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 2</TitleRoom>
                        <PlayersRoom>7/10 Jogadores | 0/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 3</TitleRoom>
                        <PlayersRoom>9/10 Jogadores | 0/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 4</TitleRoom>
                        <PlayersRoom>9/10 Jogadores | 1/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 5</TitleRoom>
                        <PlayersRoom>6/10 Jogadores | 0/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 6</TitleRoom>
                        <PlayersRoom>2/10 Jogadores | 0/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>
                
                <ViewRoom>
                    <ViewTitles>
                        <TitleRoom>Table 7</TitleRoom>
                        <PlayersRoom>6/10 Jogadores | 0/1 Mestre</PlayersRoom>
                    </ViewTitles>
                    
                    <ViewButtonRoom><ButtonRoom>
                        <TextButtonRoom>Join</TextButtonRoom>
                    </ButtonRoom></ViewButtonRoom>
                </ViewRoom>

            </ScrollView>
        </ViewOpenRoom>

        </ImageBackground>
      </ContainerHome> 

)  
}  
}

const styles = StyleSheet.create({
  backgroundImage: {
      height: '100%',
      width: '100%',
      resizeMode: 'cover',
      alignItems: 'center',
},
  container: {
    flex: 1,
    backgroundColor: '#D9BA8E',
    alignItems: 'center',
    justifyContent: 'center',
  },
  body: {
      backgroundColor: '#D9BA8E'
  }
});

























/* export default class Home extends Component{  
    render() {  
        return(  

    <Container>
        <Card>
            <Image source={require('../../assets/perfil_image.png')} />
            <Info>
                <Nome>Nome do Jogador</Nome>
                <InfoProfile>Aqui é onde vai a definição do jogador, onde informa o que ele joga.</InfoProfile>
            </Info>
        </Card>
        <ButtonStyles>
            <Button color={"#fbae5c"}><IconDeslike source={require('../../assets/tresh_ico_withe.png')} /></Button>
            <Button color={"#c48eff"}><IconLike source={require('../../assets/controll_ico_white.png')} /></Button>
        </ButtonStyles>
    </Container>
)  
    }  
}   */
 

/*   <Container>
            <Card>
                <Image source={require('../../assets/logo.png')}/>
                <Info>
                    <Nome>JohnnBZ</Nome>
                    <Info>Jogo LOL e Dota, sou gay duas vezes;</Info>
                    <ButtonStyles>
                    <Button><Icon source={require('../../assets/logo.png')}/></Button>
                    <Button><Icon source={require('../../assets/logo.png')}/></Button>
                    </ButtonStyles>
                </Info>
            </Card>
       </Container>      */