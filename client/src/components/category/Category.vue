<template>
    <div class="row">
        <h1>Categories</h1>
        <hr>
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        methods: {
            navigateToHome() {
                this.$router.push({ name: 'home' });
            }
        },
        computed: {
            
        }
    }
</script>