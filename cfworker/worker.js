addEventListener('fetch', event => {
  event.respondWith(handleRequest(event.request))
})

async function handleRequest(request) {
  if (request.method !== 'GET') {
    return new Response('Method not allowed', { status: 405 })
  }

  const url = new URL(request.url)
  
  const pathMatch = url.pathname.match(/\/ip\/(.+)/)
  if (!pathMatch) {
    return new Response('Invalid path. Use /ip/<IP_ADDRESS>', { status: 400 })
  }

  const ip = pathMatch[1]
  const targetUrl = 'https://scamalytics.com/ip/' + ip

  const headers = new Headers(request.headers)
  
  ['x-forwarded-for', 'x-real-ip', 'forwarded'].forEach(h => headers.delete(h))
  
  const customHeaders = {
    'User-Agent': 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36',
    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
    'Accept-Language': 'en-US,en;q=0.5',
    'Accept-Encoding': 'gzip, deflate, br, zstd',
    'Sec-GPC': '1',
    'Connection': 'keep-alive',
    'Priority': 'u=0, i',
    'Pragma': 'no-cache',
    'Cache-Control': 'no-cache',
  }
  
  Object.entries(customHeaders).forEach(([key, value]) => headers.set(key, value))

  try {
    const response = await fetch(targetUrl, { 
      method: 'GET',
      headers: headers 
    })
    
    const resHeaders = new Headers(response.headers)
    
    resHeaders.set('Access-Control-Allow-Origin', '*')
    resHeaders.set('Access-Control-Allow-Credentials', 'true')
    
    ['Content-Security-Policy', 'Content-Security-Policy-Report-Only', 'Clear-Site-Data'].forEach(h => {
      resHeaders.delete(h)
    })
    
    return new Response(response.body, { 
      status: response.status, 
      headers: resHeaders 
    })
  } catch (error) {
    return new Response('Error fetching from Scamalytics: ' + error.message, { 
      status: 500 
    })
  }
}

