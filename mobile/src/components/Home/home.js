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
    
    /* constructor(props) {
        super(props);
        this.state = {
            data: []
        }
    }

    componentDidMount() {
        fetch('https://randomuser.me/api/?results=10')
            .then( res => res.json() )
            .then( res => {
                this.setState({
                    data: res.results || []
                })
                console.log(this.state.data)
    })} */

    state = {
        nameList: [],
        firstName: []
    }

    componentDidMount() {
        axios.get('https://randomuser.me/api/?results=10')
        
            .then(res => {
                const nameList = res.data.results;
                const firstName = res.data.results;
                this.setState({ nameList });
                this.setState({ firstName });
                console.log(this.state.firstName[0].name.first)
            })
            
      }

    render() {

        const {firstName} = this.state;
        const {nameList} = this.state;

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
                            {this.state.firstName.map(user => <TextBoxNameProfile>Bem Vindo {user.name.first}</TextBoxNameProfile>)}
                        </ViewTextProfile>
                    </ViewHeaderHome>

                    <TitleHome>Salas Abertas</TitleHome>
                    <Input placeholder="Pesquisar salas.." />

                    <ViewOpenRoom>
                        <ScrollView>

                              <FlatList
                                data={this.state.nameList}
                                renderItem={({ item }) =>

                                    <ViewRoom>
                                        <ViewTitles>
                                            <TitleRoom>{item.name.first}</TitleRoom>
                                            <PlayersRoom>/10 Jogadores | /1 Mestre</PlayersRoom>
                                        </ViewTitles>

                                        <ViewButtonRoom><ButtonRoom>
                                            <TextButtonRoom>Join</TextButtonRoom>
                                        </ButtonRoom></ViewButtonRoom>
                                    </ViewRoom>

                                } />  

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