import { useState, useEffect } from 'react'
import './App.css'
import ProfessionalCard from './components/ProfessionalCard'
import { FaPaw, FaTooth, FaAmbulance } from 'react-icons/fa'

function App() {
  const [professionals, setProfessionals] = useState([])
  const [selectedCategory, setSelectedCategory] = useState(null)

  useEffect(() => {
    if (selectedCategory === null) return
    async function fetchProfessionals() {
      const response = await fetch(`http://localhost:8080/professionals?category=${selectedCategory}`)
      const data = await response.json()
      setProfessionals(data)
    }
    fetchProfessionals()
  }, [selectedCategory])

  const categories = [
    { value: "vet", label: "Veterinarias", icon: FaPaw },
    { value: "dentist", label: "Odontología", icon: FaTooth },
    { value: "emergencies", label: "Emergencias", icon: FaAmbulance },
  ]

  return (
    <div>
      <h1>Ushuaia 24hs</h1>
      <p className="subtitle">Guardias y profesionales disponibles</p>

      {selectedCategory === null ? (
        <div>
          {categories.map((cat) => {
            const Icon = cat.icon
            return (

              <button
                key={cat.value}
                className="category-button"
                onClick={() => setSelectedCategory(cat.value)}
              >
                <Icon className="category-icon" />
                {cat.label}
              </button>
            )
          })}

        </div>
      ) : (
        <div>
          {professionals.map((prof) => (
            <ProfessionalCard key={prof.id} professional={prof} />
          ))}
          <button
            className="back-button"
            onClick={() => setSelectedCategory(null)}
          >
            ← Volver
          </button>
        </div>
      )}
    </div>
  )
}

export default App