<template>
    <md-list-item>
        <div class="md-list-item-text">
        <span>
            <router-link
                tag="li"
                :to="'/products/' + this.item.ID"
                class="list-group-item"
                style="cursor: pointer"><a href="javascript:void(0)">{{ this.item.Name }}</a></router-link>
        </span>
        <span>id: {{ this.item.ID }}</span>
        <span>Price: {{ this.item.Price }}</span>
        </div>
    </md-list-item>
</template>

<script>
export default {
    props: ['item'],
}
</script>