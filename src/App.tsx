import { useEffect, useState } from 'react'

// items in inventory contain these attributes
type ItemType = {
  name: string
  description: string
  location: string
}

// url of lambda function which will act as our backend
const endpoint: string = "https://elmadre:8000/api"

function App() {
  // list of all items in our inventory
  const [inventory, setInventory] = useState<any>();

  // retrive inventory information from database
  useEffect(() => {
    fetch(endpoint)
      .then(res => res.json())
      .then(data => setInventory(data))
      .catch(e => console.error(e))
  }, [])

  return (
    <div>
    </div>
  );
}

export default App;
