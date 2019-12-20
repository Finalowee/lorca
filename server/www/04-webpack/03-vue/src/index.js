const {add, mul} = require('./js/tool.js')

// import css from './css/normal.css'
import less from './css/special.less'

import Vue from 'vue'
import App from './components/app.vue'

new Vue({
    el: '#app',
    template: '<App/>',
    components: {
        App
    }
})
