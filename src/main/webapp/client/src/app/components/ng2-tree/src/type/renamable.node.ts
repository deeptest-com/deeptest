
export interface RenamableNode {
  /**
   * Set new value of the renamable node. Implementation of this method is up to user.
   * @param {string} name - A new value of the node.
   */
  setName(name: string): void;

  /**
   * Get string representation of the node. Implementation of this method is up to user.
   * @returns {string} - A node string representation.
   */
  toString(): string;
}
