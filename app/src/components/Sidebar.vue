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
    <div class="title-box">
      <h1>Flowstate</h1>
    </div>
    <div class="instruction-text">Drag nodes to the graph</div>
    
    <div class="nodes-table">
      <div 
        v-for="node in nodeTypes" 
        :key="node"
        :class="['node-row', `vue-flow__node-${node}`]"
        :draggable="true" 
        @dragstart="onDragStart($event, node)"
      >
        <div class="wrapper">
          <div class="label">{{ nodeTypesMap[node] }}</div>
        </div>
      </div>
    </div>
  </aside>
</template>

<style scoped>
@import '../assets/nodes.css';
@import '../assets/main.css';

aside {
  position: fixed;
  left: 0;
  top: 0;
  height: 100vh;
  transform: translateX(0);
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
  transform: translateX(-100%);
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

.nodes-table {
  display: flex;
  flex-direction: column;
  gap: 10px;
  width: 180px;
  margin-top: 20px;
}

.node-row {
  padding: 12px;
  border-radius: 6px;
  margin-bottom: 8px;
  height: 50px;
  width: 100%;
  display: flex;
  align-items: center;
}

.wrapper {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.vue-flow__node {
  cursor: grab;
  width: 100%;
  height: 100%;
  min-width: unset;
  max-width: 100%;
  min-height: unset;
  transition: all 0.2s ease;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #222222;
  border: 1px solid rgba(255, 255, 255, 0.12);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
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

.title-box {
  width: 100%;
  padding: 15px 0;
  text-align: center;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  margin-bottom: 20px;
}

.title-box h1 {
  font-size: 32px;
  font-weight: 600;
  color: #ffffff;
  letter-spacing: 0.5px;
  margin: 0;
  text-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.instruction-text {
  text-align: center;
  color: rgba(255, 255, 255, 0.7);
  font-size: 0.9rem;
  margin: 10px 0;
}
</style>
