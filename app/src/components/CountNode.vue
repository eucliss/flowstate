<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'  
import { Position, Handle } from '@vue-flow/core'
import { computed } from 'vue'
import { getNodeStatus } from '../functions'
import { ref } from 'vue'
// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<NodeProps>()
const emit = defineEmits(['updateNodeInternals'])
const haveHandles = ref(false)

// Replace ref and onMounted with a computed property
const count = ref(0)
count.value = 20

</script>

<template>
    <div v-if="haveHandles">
        <Handle class="vue-flow__handle horizontal" id="bottom" type="source" :position="Position.Bottom" />
        <Handle class="vue-flow__handle horizontal" id="top" type="target" :position="Position.Top" />
        <Handle class="vue-flow__handle vertical" id="left" type="source" :position="Position.Left" />
        <Handle class="vue-flow__handle vertical" id="right" type="source" :position="Position.Right" />    
    </div>
    <div 
        :class="['vue-flow__node']"
        style="position: relative; z-index: 1;"
    >
        <div class="wrapper">
            <div class="label">{{ props.data.label }}</div>
            <div class="count">{{ count }}</div>
        </div>
    </div>
</template>

<style scoped>
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';

/* import the default theme, this is optional but generally recommended */
@import '@vue-flow/core/dist/theme-default.css';

.wrapper {
    position: relative;
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
}

.label {
    position: absolute;
    top: 4px;
    left: 4px;
    font-size: 10px;
    color: #7d7d7d;
}

.count {
    font-size: 30px;
    color: #ffffff;
}

.vue-flow__node {
  background-color: #000000;
  position: relative;
  z-index: 1;
  min-width: 150px;
  min-height: 40px;
  height: 60px;
  width: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid white;
  border-radius: 4px;
}
.vue-flow__handle.horizontal{
    background-color:#000000;
    height:13px;
    width:50px;
    border-radius:2px;
}

.vue-flow__handle.vertical{
    background-color:#000000;
    height:24px;
    width:12px;
    border-radius:2px;
}
</style>
