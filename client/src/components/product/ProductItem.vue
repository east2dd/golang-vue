<template>
    <div>
        <router-link
            tag="p"
            :to="'/products/' + this.item.ID"
            style="cursor: pointer"><a href="javascript:void(0)">{{ this.item.Name }}</a></router-link>

        <p>Price: {{ this.item.Price }}</p>
        <p>{{ this.item.Description }}</p>
    </div>
</template>

<script>
export default {
    props: ['item'],
}
</script>