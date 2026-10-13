import { useState, useEffect } from 'react'
import './App.css'
import ProfessionalCard from './components/ProfessionalCard'

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
  { value: "vet", label: "Veterinarias" },
  { value: "dentist", label: "Odontología" },
  { value: "emergencies", label: "Emergencias" },
]

return (
  <div>
    <h1>Ushuaia 24hs</h1>
    <p>Guardias y profesionales disponibles</p>

    {selectedCategory === null ? (
      <div>
        {categories.map((cat) => (
          <button
          key={cat.value}
          className="category-button"
          onClick={() => setSelectedCategory(cat.value)}
          >
            {cat.label}
          </button>
        ))}
      </div>
    ): (
      <div>
        <button
        className="back-button"
        onClick={() => setSelectedCategory(null)}
        >
          Volver
        </button>
        {professionals.map((prof) => (
          <ProfessionalCard key={prof.id} professional={prof} />
        ))}
      </div>
    )}
  </div>
)
}

export default App