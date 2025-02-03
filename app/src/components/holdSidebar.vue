<script setup>
import useDragAndDrop from '../useDnD'
import { nodeTypesMap, nodeTypes } from '../functions'

const { onDragStart } = useDragAndDrop()
</script>

<template>
  <aside class="sidebar">
    <div class="nodes" v-for="node in nodeTypes" :key="node">
      <div 
        :class="['vue-flow__node', `vue-flow__node-${node}`]"
        :draggable="true" 
        @dragstart="onDragStart($event, node)"
      >
        <div class="wrapper">
          <div class="label">{{ nodeTypesMap[node] }}</div>
        </div>
      </div>

      <div class="nodes">
        <template v-for="node in nodeTypes" :key="node">
          
        </template>
      </div>
    </div>
  </aside>
</template>

<style scoped>
@import '../assets/nodes.css';

aside {
  position: fixed;
  left: 0;
  top: 50%;
  transform: translateY(-50%);
  padding: 10px;
  z-index: 999;
  pointer-events: auto;
}

.nodes {
  margin-bottom: 20px;  /* Increased margin for better separation */
  display: block;       /* Ensures each node container is on its own line */
  width: 100%;         /* Takes full width of the sidebar */
}

/* Override some base node styles for sidebar */
.vue-flow__node {
  cursor: grab;
  width: 140px;  /* Fixed width for sidebar nodes */
  height: 50px;  /* Fixed height for sidebar nodes */
  min-width: unset;
  min-height: unset;
  transition: all 0.2s ease;
  position: relative;    /* Ensures proper stacking */
  display: block;       /* Prevents inline behavior */
  margin: 0 auto;       /* Centers the node in the container */
}

.vue-flow__node:hover {
  transform: scale(1.05);
  box-shadow: 0 0 10px rgba(255, 255, 255, 0.2);
}
</style>
