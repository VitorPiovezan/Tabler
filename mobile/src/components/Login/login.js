import React, { Component } from 'react';
import { StatusBar, AsyncStorage, StyleSheet, ImageBackground} from 'react-native';
import axios from 'axios';

import {
    Logo, 
    Container, 
    SubTitle, 
    Title, 
    Inputs,
    ButtonView,
    Button,
    TextButton,
    ErrorMessage
} from './styles';
import api from '../../api/api';


export default class Login extends Component{  

    state = {
        email: 'tabler@tabler.com',
        password: 'tabler',
        error: '',
    }

    handleNameChange = (email) => {
        this.setState({ email });
        console.log(this.state.email);
    }

    handlePasswordChange = (password) => {
        this.setState({ password });
        console.log(this.state);
    }

    handleSignupPress = () => {
        this.props.navigation.navigate('Signup');
    }

    handleSignInPress = async () => {
        this.props.navigation.navigate('Home');
            
}

    render(){  
        console.disableYellowBox = true; //para ignorar warnigns
        return( 
            
            <Container>
              <ImageBackground source={require('../../assets/fundo.png')} style={styles.backgroundImage}>  
            <StatusBar
                barStyle="light-content"
                backgroundColor="#5E3200"
            />
                <Logo source={require('../../assets/logo.png')}/>
                <Title>Login</Title>
                
                <Inputs 
                    autoCapitalize='none'
                    onChangeText={this.handleNameChange}
                    value={this.state.email}
                    placeholder="Digite seu e-mail" 
                    keyboardType="email-address" 
                    returnKeyType = {"next"}/>

                <Inputs
                    autoCapitalize='none'
                    onChangeText={this.handlePasswordChange} 
                    value={this.state.password}
                    placeholder="Digite sua senha" 
                    secureTextEntry={true}/>
                
                {this.state.error.length !== 0 && <ErrorMessage>{this.state.error}</ErrorMessage>}

                <ButtonView>
                    <Button onPress={this.handleSignInPress}>
                        <TextButton>Login</TextButton>
                    </Button>
                    <Button
                        onPress={this.handleSignupPress}
                    >
                        <TextButton>Cadastre-se</TextButton>
                    </Button>
                </ButtonView>
                </ImageBackground>
            </Container>
        )}}

        const styles = StyleSheet.create({
            backgroundImage: {
                height: '100%',
                width: '100%',
                alignItems: 'center',
                resizeMode: 'cover',
                justifyContent: 'center',
          }
          });

Login.navigationOptions = {
    header: null
  }

  