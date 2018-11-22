<template>
    <div>
        <h1>The User Page</h1>
        <p><button @click="navigateToHome" class="btn btn-primary">Go to Home</button></p>
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