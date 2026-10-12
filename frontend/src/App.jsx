import { useState, useEffect } from 'react'
import './App.css'

function App() {
  const [professionals, setProfessionals] = useState([])

  useEffect(() => {
    async function fetchProfessionals() {
      const response = await fetch("http://localhost:8080/professionals")
      const data = await response.json()
      setProfessionals(data)
      console.log(data)
    }
    fetchProfessionals()
  }, [])

return (
  <div>
    <h1>Ushuaia 24hs</h1>
    <p>Guardias y profesionales disponibles</p>

    {professionals.map((prof) => (
      <div key={prof.id} className="card">
        <h2>{prof.name}</h2>
        <p>{prof.category}</p>
        <p>{prof.address}</p>
        <a href={`tel:${prof.phone}`} className="call-button">
          Llamar
        </a>
        <a href={`https://www.google.com/maps/search/?api=1&query=${encodeURIComponent(prof.address + ", Ushuaia")}`}
        target="_blank"
        rel="noopener noreferrer"
        className="map-button">
          Cómo llegar
        </a>
      </div>
    ))}

  </div>
)
}

export default App