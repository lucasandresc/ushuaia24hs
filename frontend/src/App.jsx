import { useState, useEffect } from 'react'
import './App.css'
import ProfessionalCard from './components/ProfessionalCard'

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

const categories = [
  { value: "vet", label: "Veterinarias" },
  { value: "dentist", label: "Odontología" },
  { value: "emergencies", label: "Emergencias" },
]

return (
  <div>
    <h1>Ushuaia 24hs</h1>
    <p>Guardias y profesionales disponibles</p>

    {professionals.map((prof) => (
      <ProfessionalCard key={prof.id} professional={prof} />
    ))}

  </div>
)
}

export default App