import { useVueFlow } from '@vue-flow/core'
import { ref, watch } from 'vue'
import { addNode, nodeTypesMap } from './functions'
import type { Node } from '@vue-flow/core'
/**
 * In a real world scenario you'd want to avoid creating refs in a global scope like this as they might not be cleaned up properly.
 * @type {{draggedType: Ref<string|null>, isDragOver: Ref<boolean>, isDragging: Ref<boolean>}}
 */
const state = {
  /**
   * The type of the node being dragged.
   */
  draggedType: ref(null),
  isDragOver: ref(false),
  isDragging: ref(false),
}

export default function useDragAndDrop() {
  const { draggedType, isDragOver, isDragging } = state

  const { addNodes, screenToFlowCoordinate, onNodesInitialized, updateNode } = useVueFlow()

  watch(isDragging, (dragging) => {
    document.body.style.userSelect = dragging ? 'none' : ''
  })

  function onDragStart(event, type) {
    if (event.dataTransfer) {
      event.dataTransfer.setData('application/vueflow', type)
      event.dataTransfer.effectAllowed = 'move'
    }

    draggedType.value = type
    isDragging.value = true

    document.addEventListener('drop', onDragEnd)
  }

  /**
   * Handles the drag over event.
   *
   * @param {DragEvent} event
   */
  function onDragOver(event) {
    event.preventDefault()

    if (draggedType.value) {
      isDragOver.value = true

      if (event.dataTransfer) {
        event.dataTransfer.dropEffect = 'move'
      }
    }
  }

  function onDragLeave() {
    isDragOver.value = false
  }

  function onDragEnd() {
    isDragging.value = false
    isDragOver.value = false
    draggedType.value = null
    document.removeEventListener('drop', onDragEnd)
  }

  /**
   * Handles the drop event.
   *
   * @param {DragEvent} event
   */
  function onDrop(event) {
    event.preventDefault()
    event.stopPropagation()
    
    if (!draggedType.value) return
    
    const position = screenToFlowCoordinate({
      x: event.clientX,
      y: event.clientY,
    })
    
    const nodeId = crypto.randomUUID()
    const new_node: Node = {
        id: nodeId,
        type: draggedType.value,
        position: position,
        data: {
            label: "New " + nodeTypesMap[draggedType.value],
            sql: "",
            successRoute: {
                leftValue: "",
                rightValue: "",
                operator: "",
            },
            failureRoute: {
                leftValue: "",
                rightValue: "",
                operator: "",
            },
            status: false,
            count: 0,
        },
    }
    addNode(new_node)

    // /**
    //  * Align node position after drop, so it's centered to the mouse
    //  *
    //  * We can hook into events even in a callback, and we can remove the event listener after it's been called.
    //  */
    // const { off } = onNodesInitialized(() => {
    //   updateNode(nodeId, (node) => ({
    //     position: { x: node.position.x - node.dimensions.width / 2, y: node.position.y - node.dimensions.height / 2 },
    //   }))

    //   off()
    // })
  }

  return {
    draggedType,
    isDragOver,
    isDragging,
    onDragStart,
    onDragLeave,
    onDragOver,
    onDrop,
  }
}
