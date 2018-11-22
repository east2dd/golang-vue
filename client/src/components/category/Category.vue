<template>
    <div>
        <h1>The Category Page</h1>
        <button @click="navigateToHome" class="btn btn-primary">Go to Home</button>
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