<script setup lang="ts">
import type { NodeProps } from '@vue-flow/core'  
import { Position, Handle } from '@vue-flow/core'
import { computed } from 'vue'
import { getNodeStatus } from '../functions'

// props were passed from the slot using `v-bind="customNodeProps"`
const props = defineProps<NodeProps>()

// Replace ref and onMounted with a computed property
const status = computed(() => props.data.status)

</script>

<template>
  <div 
    :class="[
      'vue-flow__node',
      status ? 'vue-flow__node-success' : 'vue-flow__node-failed'
    ]"
    style="position: relative;"
  >
    <Handle type="target" :position="Position.Top" :connectable="true" />
    <div>{{ props.data.label }}</div>
    <Handle type="source" :position="Position.Bottom" :connectable="true" />
  </div>
</template>

<style scoped>
/* import the necessary styles for Vue Flow to work */
@import '@vue-flow/core/dist/style.css';

/* import the default theme, this is optional but generally recommended */
@import '@vue-flow/core/dist/theme-default.css';

.vue-flow__node-success,
.vue-flow__node-failed {
  min-width: 150px;
  min-height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.vue-flow__node-success {
  background: green;
  color: white;
  border: 1px solid green;
  border-radius: 4px;
  box-shadow: 0 0 0 1px green;
  padding: 8px;
}

.vue-flow__node-failed {
  background: red;
  color: white;
  border: 1px solid red;
  border-radius: 4px;
  box-shadow: 0 0 0 1px red;
  padding: 8px;
}
</style>
