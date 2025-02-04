<script setup>
import { Background } from '@vue-flow/background'
</script>

<template>
    <div class="dropzone-background">
        <Background 
            class="dots-background"
            variant="dots"
            height="100%"
            width="100%"
            size="2"
            gap="25"
        />
        <div class="overlay">
            <slot />
        </div>
    </div>
</template>

<style>
@import '@vue-flow/core/dist/style.css';
@import '@vue-flow/core/dist/theme-default.css';
@import '@vue-flow/controls/dist/style.css';

.dropzone-background {
    position:relative;
    height:100%;
    width:100%
}

.dropzone-background .overlay {
    position:absolute;
    top:0;
    left:0;
    height:100%;
    width:100%;
    display:flex;
    align-items:center;
    justify-content:center;
    z-index:1;
    pointer-events:none
}

</style>
