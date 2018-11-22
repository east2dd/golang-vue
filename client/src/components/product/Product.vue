<template>
    <div>
        <h1>Products</h1>
        <router-view></router-view>
    </div>
</template>

<script>
    export default {
        methods: {
            navigateToHome() {
                this.$router.push({ name: 'home' });
            }
        }
    }
</script>