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

import { ScrollView, FlatList } from 'react-native-gesture-handler';
import { StyleSheet, ImageBackground, Text } from 'react-native';
import axios from 'axios';
export default class Home extends Component {
    
    constructor(props) {
        super(props);
        this.state = {
            users: []
    }}

    componentDidMount() {
        axios.get('https://jsonplaceholder.typicode.com/users')   
        .then(res => {
                this.setState({ users: res.data });
                console.log(this.state);
            })
      }

    render() {

        return (
            <ContainerHome>
                <ImageBackground source={require('../../assets/fundo.png')} style={styles.backgroundImage} >

                    <ViewHeaderHome>
                        <ViewConfigsProfile>
                            <ButtonProfile onPress={() => this.props.navigation.navigate('Config')}>
                                <IconsProfile source={require('../../assets/configIcon.png')} />
                            </ButtonProfile>
                            <ImgHomeConfig source={require('../../assets/icon.png')} />
                            <ButtonProfile onPress={() => this.props.navigation.navigate('EditProfile')} >
                                <IconsProfile source={require('../../assets/perfilIcon.png')} />
                            </ButtonProfile>
                        </ViewConfigsProfile>

                        <ViewTextProfile>
                            <TextBoxNameProfile>Bem Vindo Caralho</TextBoxNameProfile>
                        </ViewTextProfile>
                    </ViewHeaderHome>

                    <TitleHome>Salas Abertas</TitleHome>
                    <Input placeholder="Pesquisar salas.." />

                    <ViewOpenRoom>
                        <ScrollView>


                                    {this.state.users.map(item => <ViewRoom>
                                        <ViewTitles>
                                            <TitleRoom>{item.name}</TitleRoom>
                                            <PlayersRoom>/10 Jogadores | /1 Mestre</PlayersRoom>
                                        </ViewTitles>

                                        <ViewButtonRoom><ButtonRoom>
                                            <TextButtonRoom>Join</TextButtonRoom>
                                        </ButtonRoom></ViewButtonRoom>
                                    </ViewRoom>)}                                

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