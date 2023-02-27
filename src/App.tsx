/**
 * Sample React Native App
 * https://github.com/facebook/react-native
 *
 * @format
 */

import React from 'react';
import type {PropsWithChildren} from 'react';
import {
  ScrollView,
  StatusBar,
  StyleSheet,
  Text,
  useColorScheme,
  View,
  ViewStyle,
} from 'react-native';

import color  from './style/Color';

import { SafeAreaProvider, SafeAreaView } from 'react-native-safe-area-context';
import Icon from 'react-native-vector-icons/FontAwesome';


function App(): JSX.Element {
  const isDarkMode = true;

  const backgroundStyle: ViewStyle = {
    flex: 1,
    // backgroundColor: color.backgroundPrimaryColor,
  };

  return (
    <SafeAreaProvider>
      <SafeAreaView
        edges={['left', 'right', 'bottom']}
        style={backgroundStyle}>
        <StatusBar
          barStyle={isDarkMode ? 'light-content' : 'dark-content'}
        />
        <ScrollView
          contentInsetAdjustmentBehavior='never'
          style={{backgroundColor: color.backgroundHighLightColor}}
          showsVerticalScrollIndicator={false}
        >
          <PageHeader/>
          <ServicesView/>
          <MessageView/>
          <BannerView/>
          <FundView/>
          <LoadView/>
        </ScrollView>
      </SafeAreaView>
    </SafeAreaProvider>
  );
}

/// Header
type PageHeaderProps = {}

function PageHeader (props: PageHeaderProps) {
  return (
      <View
        style={{
          backgroundColor: 'transparent',
          height: 150,
          paddingBottom: 10,
          flexDirection: 'row',
          alignItems: 'flex-end',
          justifyContent: 'space-between',
          paddingHorizontal: 20
        }}>
      <HeaderItem icon='money' title='朝朝宝'/>
      <HeaderItem icon='bitcoin' title='收支明细'/>
      <HeaderItem icon='arrows-h' title='转帐'/>
      <HeaderItem icon='file-text' title='账户预览'/>
      </View>
  );
}

function HeaderItem (props: {icon: string, title: string}) {
  return (
    <View style={{ justifyContent: 'center', alignItems: 'center' }}>
      <Icon name={props.icon} size={30} color='white'/>
      <Text style={{ fontSize: 18, fontWeight: '600', color: 'white' }}>{props.title}</Text>
    </View>
  )
}


/// 服务
type ServicesViewProps = {}

function ServicesView(props: ServicesViewProps) {
  return (
    <View
      style={{
        backgroundColor: color.backgroundPrimaryColor,
        borderTopLeftRadius: 20,
        borderTopRightRadius: 20,
        height: 150
      }}>
    </View>
  );
}


/// 消息跑马灯
type MessageViewProps = {}

function MessageView(props: MessageViewProps) {
  return (
    <View
      style={{
        backgroundColor: color.backgroundPrimaryColor,
        height: 50 
      }}>
    </View>
  );
}


/// Banner
type BannerViewProps = {}

function BannerView(props: BannerViewProps) {
  return (
    <View
      style={{
        backgroundColor: color.backgroundPrimaryColor,
        height: 100
      }}
    ></View>
  );
}


/// 财富精选
type FundViewProps = {}

function FundView(props: FundViewProps) {
  return (
    <View
      style={{
        backgroundColor: color.backgroundPrimaryColor,
        height: 250
      }}
    ></View>
  );
}

/// 贷款产品
type LoadViewProps = {}

function LoadView(props: LoadViewProps) {
  return (
    <View
      style={{
        height: 250,
        backgroundColor: color.backgroundPrimaryColor
      }}
    ></View>
  );
}

const styles = StyleSheet.create({
  sectionContainer: {
    marginTop: 32,
    paddingHorizontal: 24,
  },
  sectionTitle: {
    fontSize: 24,
    fontWeight: '600',
  },
  sectionDescription: {
    marginTop: 8,
    fontSize: 18,
    fontWeight: '400',
  },
  highlight: {
    fontWeight: '700',
  },
});

export default App;
