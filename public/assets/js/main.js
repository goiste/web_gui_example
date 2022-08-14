new Vue({
  el: '#app',
  vuetify: new Vuetify(),
  data () {
    return {
      host: {},
      cpu: {},
      ram: {},
    }
  },
  filters: {
    datetime (val) {
      if (!val) return
      return new Date(Date.parse(val)).toLocaleString()
    }
  },
  methods: {
    loadData () {
      this.loadHost()
      this.loadCpu()
      this.loadRam()
    },
    loadHost () {
      fetch('/api/v1/host').then(res => res.json()).then(res => this.host = Object.assign({}, res.host))
    },
    loadCpu () {
      fetch('/api/v1/cpu').then(res => res.json()).then(res => this.cpu = Object.assign({}, res.cpu))
    },
    loadRam () {
      fetch('/api/v1/mem').then(res => res.json()).then(res => this.ram = Object.assign({}, res.mem))
    },
  },
  mounted () {
    const loc = window.location
    const ws = new WebSocket(`ws://${loc.host + loc.pathname}ws`)

    ws.onopen = function () {
      console.log('socket connected')
      setInterval(function ping () {
        ws.send('')
        return ping
      }(), 1000)
    }

    ws.onmessage = evt => {
      const upd = JSON.parse(evt.data)

      this.host.uptime = upd.uptime
      this.host.processes = upd.processes

      this.cpu.used_percent = upd.cpu_used_percent

      this.ram.used = upd.mem_used
      this.ram.free = upd.mem_free
      this.ram.used_percent = upd.mem_used_percent
    }

    this.loadData()
  }
})