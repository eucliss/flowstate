<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'  
import { Position, Handle } from '@vue-flow/core'
import { computed } from 'vue'
import { getNodeStatus } from '../../functions'
import { ref } from 'vue'
// Remove SVG import and add PNG import
import flow2 from '../../assets/flow2.png'
// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<NodeProps>()
const emit = defineEmits(['updateNodeInternals'])


</script>

<template>
    <div class="custom-node-container nodrag">
        <Handle class="vue-flow__handle" id="bottom" type="source" :position="Position.Bottom" />
        <div class="flow-node">
            <img :src="flow2" class="flow-icon" alt="Flow icon" />
        </div>
    </div>
</template>

<style scoped >
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';

/* import the default theme, this is optional but generally recommended */
@import '@vue-flow/core/dist/theme-default.css';

.vue-flow__node-custom {
    padding: 10px;
    border-radius: 5px;
    position: relative;
}

.flow-node {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: white;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    padding: 2px;
    box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
    overflow: hidden;
}

.flow-icon {
    width: 100%;
    height: 100%;
    object-fit: contain;
    padding: 1px;
}

/* Remove or update the deep selector since we're using a class */
/* .flow-node :deep(svg) {
    width: 100%;
    height: 100%;
} */
</style>
