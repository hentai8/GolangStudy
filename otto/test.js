function decodeDepositEventEventData(eventData) {
    // import {onchain_events} from '@starcoin/starcoin'
    var {onchain_events} = require('@starcoin/starcoin');
    var eventName = 'DepositEvent';
    return onchain_events.decodeEventData(eventName, eventData).toJS();
}

const {onchain_events} = require('@starcoin/starcoin')
var eventName = 'DepositEvent';
var eventData = '0x00f2052a01000000000000000000000000000000000000000000000000000001035354430353544300';

console.log(onchain_events.decodeEventData(eventName, eventData).toJS());