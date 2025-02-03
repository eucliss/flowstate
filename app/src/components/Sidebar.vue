<script setup>
import { ref } from 'vue'
import useDragAndDrop from '../useDnD'
import { nodeTypesMap, nodeTypes } from '../functions'

const { onDragStart } = useDragAndDrop()
const isOpen = ref(true)

const toggleSidebar = () => {
  isOpen.value = !isOpen.value
}
</script>

<template>
  <aside :class="{ 'closed': !isOpen }">
    <button class="toggle-button" @click="toggleSidebar">
      <span class="carrot" :class="{ 'rotated': isOpen }">›</span>
    </button>
      <div 
        v-for="node in nodeTypes" :key="node"
        :class="[`vue-flow__node-${node}`]"
        :draggable="true" 
        @dragstart="onDragStart($event, node)"
      >
        <div class="wrapper">
          <div class="label">{{ nodeTypesMap[node] }}</div>
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
  padding: 20px;
  z-index: 999;
  pointer-events: auto;
  background: #181818;
  border-right: 1px solid rgba(255, 255, 255, 0.1);
  box-shadow: 
    0 0 30px rgba(0, 0, 0, 0.5),
    inset 0 0 20px rgba(0, 0, 0, 0.5);
  border-radius: 0 15px 15px 0;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
}

aside.closed {
  transform: translate(-100%, -50%);
}

.toggle-button {
  position: absolute;
  right: -30px;
  top: 50%;
  transform: translateY(-50%);
  width: 30px;
  height: 60px;
  background: #181818;
  border: 1px solid rgba(255, 255, 255, 0.1);
  border-left: none;
  border-radius: 0 8px 8px 0;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  box-shadow: 5px 0 15px rgba(0, 0, 0, 0.5);
  transition: all 0.3s ease;
}

.toggle-button:hover {
  background: rgba(15, 15, 20, 0.95);
  box-shadow: 
    5px 0 20px rgba(0, 0, 0, 0.3),
    inset 0 0 15px rgba(255, 255, 255, 0.05);
}

.carrot {
  color: rgba(255, 255, 255, 0.7);
  font-size: 1.5rem;
  font-weight: bold;
  transition: transform 0.3s ease;
}

.carrot.rotated {
  transform: rotate(180deg);
}

.toggle-button:hover .carrot {
  color: rgba(255, 255, 255, 0.9);
}

.nodes {
  margin-bottom: 25px;
  display: block;
  width: 100%;
  position: relative;
}

.nodes::after {
  content: '';
  position: absolute;
  bottom: -12px;
  left: 0;
  width: 100%;
  height: 1px;
  background: rgba(255, 255, 255, 0.08);
}

.vue-flow__node {
  cursor: grab;
  width: 140px;
  height: 50px;
  min-width: unset;
  min-height: unset;
  transition: all 0.2s ease;
  position: relative;
  display: block;
  margin: 0 auto;
  background: #222222;
  border: 1px solid rgba(255, 255, 255, 0.12);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
  color: white;
  font-size: 1.1rem;
  font-weight: 500;
}

.vue-flow__node:hover {
  transform: translateY(-2px);
  background: #2a2a2a;
  border-color: rgba(255, 255, 255, 0.2);
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
}

/* Remove gradient overlay */
.vue-flow__node::before {
  display: none;
}

/* Remove pulse animation */
.vue-flow__node:hover::after {
  display: none;
}

/* Target any nested elements that might contain text */
.vue-flow__node :deep(*) {
  color: white;
  font-size: 1.1rem;
  font-weight: 500;
}

/* Make sure labels are centered and properly sized */
.vue-flow__node :deep(.label) {
  color: white;
  font-size: 1.1rem;
  text-align: center;
  width: 100%;
  padding: 0 10px;
}
</style>
