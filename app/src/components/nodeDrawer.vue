<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps({
  selectedNode: {
    type: Object,
    default: () => null
  }
})

const emit = defineEmits(['close', 'addNode', 'updateNode'])
const nodeName = ref('')
const sqlQuery = ref('')

// Watch for changes to selectedNode and update form
watch(() => props.selectedNode, (newNode) => {
  if (newNode) {
    nodeName.value = newNode.name
    sqlQuery.value = newNode.sqlQuery
  } else {
    nodeName.value = ''
    sqlQuery.value = ''
  }
}, { immediate: true })

const handleSubmit = () => {
  const nodeData = { 
    name: nodeName.value,
    sqlQuery: sqlQuery.value 
  }
  
  if (props.selectedNode) {
    // If editing existing node, emit update event
    emit('updateNode', { ...props.selectedNode, ...nodeData })
  } else {
    // If creating new node, emit add event
    emit('addNode', nodeData)
  }
  
  nodeName.value = ''
  sqlQuery.value = ''
  emit('close')
}
</script>

<template>
    <Transition name="drawer">
      <div class="side-drawer">
        <div class="drawer-content">
          <div class="drawer-header">
            <h3>{{ props.selectedNode ? 'Edit Node' : 'Add New Node' }}</h3>
            <button class="close-button" @click="$emit('close')">×</button>
          </div>
          
          <div class="drawer-body">
            <div class="input-group">
              <label for="nodeName">Node Name</label>
              <input 
                id="nodeName"
                v-model="nodeName" 
                placeholder="Enter node name"
                type="text"
              />
            </div>
  
            <div class="input-group">
              <label for="sqlQuery">SQL Query</label>
              <textarea
                id="sqlQuery"
                v-model="sqlQuery"
                placeholder="Enter SQL query"
                class="sql-input"
                rows="1"
              ></textarea>
            </div>
  
            <div class="columns-container">
              <div class="column">
                <h4>Column 1</h4>
                <div class="column-content">
                  <!-- Column 1 content -->
                </div>
              </div>
              
              <div class="column">
                <h4>Column 2</h4>
                <div class="column-content">
                  <!-- Column 2 content -->
                </div>
              </div>
              
              <div class="column">
                <h4>Column 3</h4>
                <div class="column-content">
                  <!-- Column 3 content -->
                </div>
              </div>
            </div>
          </div>
  
          <div class="drawer-footer">
            <button class="btn-primary" @click="handleSubmit">
              {{ props.selectedNode ? 'Update Node' : 'Add Node' }}
            </button>
            <button class="btn-secondary" @click="$emit('close')">Cancel</button>
          </div>
        </div>
      </div>
    </Transition>
  </template>

<style scoped>
/* Base drawer styles */
.side-drawer {
  position: fixed;
  top: 0;
  right: 0;
  width: 40%;
  height: 100vh;
  background: #1e1e1e;
  border-left: 1px solid #333;
  box-shadow: -8px 0 32px rgba(0, 0, 0, 0.2);
  z-index: 1000;
}

.drawer-content {
  display: flex;
  flex-direction: column;
  height: 100%;
}

.drawer-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 24px;
  border-bottom: 1px solid #333;
}

.drawer-header h3 {
  color: #fff;
  margin: 0;
  font-size: 1.5rem;
  font-weight: 500;
}

/* Body and columns styling */
.drawer-body {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow-y: auto;
  padding: 24px;
}

.input-group {
  margin-bottom: 24px;
}

.columns-container {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
  flex: 1;
  min-height: 0;
}

.column {
  background: #2d2d2d;
  border-radius: 8px;
  padding: 16px;
  border: 1px solid #404040;
  display: flex;
  flex-direction: column;
  height: 100%;
}

.column h4 {
  color: #fff;
  margin: 0 0 16px 0;
  font-size: 1rem;
  font-weight: 500;
}

.column-content {
  flex: 1;
}

/* Input styling */
.input-group label {
  display: block;
  color: #999;
  margin-bottom: 8px;
  font-size: 0.9rem;
}

input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
}

input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

input::placeholder {
  color: #666;
}

.sql-input {
  width: 100%;
  padding: 12px 16px;
  background: #2d2d2d;
  border: 1px solid #404040;
  border-radius: 6px;
  color: #fff;
  font-size: 1rem;
  transition: all 0.2s ease;
  resize: vertical;
  min-height: 150px;
  font-family: monospace;
}

.sql-input:focus {
  outline: none;
  border-color: #0066cc;
  box-shadow: 0 0 0 2px rgba(0, 102, 204, 0.2);
}

.sql-input::placeholder {
  color: #666;
}

/* Footer styling */
.drawer-footer {
  padding: 24px;
  border-top: 1px solid #333;
  display: flex;
  justify-content: flex-end;
  gap: 12px;
}

/* Button styles */
.close-button {
  background: transparent;
  border: none;
  color: #666;
  font-size: 24px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s ease;
}

.close-button:hover {
  color: #fff;
  background: rgba(255, 255, 255, 0.1);
}

.btn-primary,
.btn-secondary {
  padding: 10px 20px;
  border-radius: 6px;
  font-size: 0.95rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary {
  background: #0066cc;
  color: white;
  border: none;
}

.btn-primary:hover {
  background: #0077ee;
}

.btn-secondary {
  background: transparent;
  color: #999;
  border: 1px solid #404040;
}

.btn-secondary:hover {
  background: rgba(255, 255, 255, 0.1);
  color: #fff;
}

/* Animation */
.drawer-enter-active,
.drawer-leave-active {
  transition: transform 0.3s ease;
}

.drawer-enter-from,
.drawer-leave-to {
  transform: translateX(100%);
}
</style> 