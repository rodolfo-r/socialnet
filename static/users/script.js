Array.from(document.querySelectorAll('.user-item a')).forEach(el => {
  el.href = `/user/${el.getAttribute('data-id')}`
})
