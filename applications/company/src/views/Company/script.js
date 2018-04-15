import debug from 'debug'

export default {
  name: 'faas-company',
  data: function () {
    return {}
  },
  beforeMount: function () {
    this.logger('beforeMount')
  },
  mounted: function () {
    this.logger('mounted')
  },
  updated: function () {
  },
  destroyed: function () {
    this.logger('destroyed')
  },
  methods: {
    logger: debug('fass:ui-company')
  }
}
