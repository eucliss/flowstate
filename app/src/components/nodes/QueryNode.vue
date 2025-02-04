<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'  
import { Position, Handle } from '@vue-flow/core'
import { computed, ref } from 'vue'
import { getNodeStatus } from '../../functions'

// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<NodeProps>()
const emit = defineEmits(['updateNodeInternals'])

// Replace ref and onMounted with a computed property
const status = computed(() => props.data.status)
const label = computed(() => props.data.label || "Query Node")

</script>

<template>
<Handle class="vue-flow__handle horizontal" id="bottom" type="source" :position="Position.Bottom" />
<Handle class="vue-flow__handle horizontal" id="top" type="target" :position="Position.Top" />
<Handle class="vue-flow__handle vertical" id="left" type="source" :position="Position.Left" />
<Handle class="vue-flow__handle vertical" id="right" type="source" :position="Position.Right" />    
  <div 
    :class="[
      'vue-flow__node',
      status ? 'vue-flow__node-success' : 'vue-flow__node-failed',
    ]"
    style="position: relative; z-index: 1;"
    >
    <div class="label">{{ label }}</div>
  </div>
</template>

<style scoped>
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';

/* import the default theme, this is optional but generally recommended */
@import '@vue-flow/core/dist/theme-default.css';

.vue-flow__node-success,
.vue-flow__node-failed {
  position: relative;
  z-index: 1;
  min-width: 150px;
  min-height: 40px;
  height: auto;
  width: 100px;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 1px solid white;
  border-radius: 4px;
  padding: 8px;
  word-wrap: break-word;
  word-break: break-word;
}

.vue-flow__node-success div,
.vue-flow__node-failed div {
  width: 100%;
  text-align: center;
  font-size: 0.9em;
  line-height: 1.2;
}

.vue-flow__node-success {
  background: green;
  color: white;
  border-radius: 4px;
  box-shadow: 0 0 0 1px green;
  padding: 8px;
}

.vue-flow__node-failed {
  background: red;
  color: white;
  border-radius: 4px;
  box-shadow: 0 0 0 1px red;
  padding: 8px;
}

.vue-flow__node-none {
  background: #222222;
  color: white;
  border-radius: 4px;
  box-shadow: 0 0 0 1px #222222;
  padding: 8px;
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
