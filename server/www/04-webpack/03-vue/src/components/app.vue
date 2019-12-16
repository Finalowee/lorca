<template id='app'>
<div>
<h1>{{message}}</h1>
</div>
</template>
<script>
    export default {
        template: '#app',
        data() {
            return {
                message: "Hello Vue"
            }
        }
    }
</script>

<style scoped>
</style>