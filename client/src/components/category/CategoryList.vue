<template>
    <div>
        <p>Please select a Category</p>
        <hr>
        <ul class="list-group">
            <router-link
                    tag="li"
                    to="/categories/1"
                    class="list-group-item"
                    style="cursor: pointer">Category 1</router-link>
            <router-link
                    tag="li"
                    to="/categories/2"
                    class="list-group-item"
                    style="cursor: pointer">Category 2</router-link>
            <router-link
                    tag="li"
                    to="/categories/3"
                    class="list-group-item"
                    style="cursor: pointer">Category 3</router-link>
        </ul>
    </div>
</template>