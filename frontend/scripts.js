const routes = [
  {
    key: 'volumeDown',
    url: '/volume/down',
  },
  {
    key: 'volumeUp',
    url: '/volume/up',
  },
  {
    key: 'volumeMute',
    url: '/volume/mute',
  },
  {
    key: 'mediaPlay',
    url: '/media/play',
  },
  {
    key: 'mediaPause',
    url: '/media/pause',
  },
  {
    key: 'mediaPrev',
    url: '/media/prev',
  },
  {
    key: 'mediaNext',
    url: '/media/next',
  },
];

const resolver = key => {
  const route = routes.find(route => route.key === key);

  fetch(route.url)
    .then(res => { console.log(res.status) })
    .catch(err => { console.error(err) })
};
