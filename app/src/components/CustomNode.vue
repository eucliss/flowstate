<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'  
import { Position, Handle } from '@vue-flow/core'

// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<NodeProps>()
  
console.log(props.id) // 'world'
</script>

<template>
  <div>
    <Handle type="target" :position="Position.Top" />
    <div>{{ props.data.label }}</div>
    <Handle type="source" :position="Position.Bottom" />
  </div>
</template>

<style>
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';

/* import the default theme, this is optional but generally recommended */
@import '@vue-flow/core/dist/theme-default.css';

.vue-flow__node-custom {  /* Changed from custom to special */
  background: red;
  color: white;
  border: 1px solid red;
  border-radius: 4px;
  box-shadow: 0 0 0 1px red;
  padding: 8px;
}

</style>
