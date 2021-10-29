const bridge = {
    async getKeys() {
        const resp = await window.backend.GoBridge.Keys();
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },
    async getContacts() {
        const resp = await window.backend.GoBridge.Contacts();
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },
    async removeKey(key) {
      const resp = await window.backend.GoBridge.RemoveKey(key);
      if(resp.error) {
          throw new Error(resp.error);
      }
      return true;
    },
    async updateKey(id, label) {
        const resp = await window.backend.GoBridge.UpdateKey(id, label);
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },
    async createKey(label) {
        const resp = await window.backend.GoBridge.CreateKey(label);
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },
    async removeContact(key) {
        const resp = await window.backend.GoBridge.RemoveContact(key);
        if(resp.error) {
            throw new Error(resp.error);
        }
        return true;
    },
    async updateContact(id, label) {
        const resp = await window.backend.GoBridge.UpdateContact(id, label);
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },
    async createContact(label, key) {
        const resp = await window.backend.GoBridge.CreateContact(label, key);
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },
    async copyToClipboard(text) {
      const resp = await window.backend.GoBridge.CopyToClipboard(text);
      if(resp.error) {
          throw new Error(resp.error);
      }
      return true;
    },
    async getSettings() {
        const resp = await window.backend.GoBridge.Settings();
        if(resp.error) {
            throw new Error(resp.error);
        }
        return resp.data;
    },

}
export default bridge;