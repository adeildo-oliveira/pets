package repository

const (
	selectPetsQuery = `SELECT id, name_pet, age_pet, breed_pet, nick_name_pet, status_pet, date_create_pet, date_update_pet FROM pets`
	insertPetQuery  = `INSERT INTO pets (name_pet, age_pet, breed_pet, nick_name_pet, status_pet, date_create_pet) 
						VALUES (:name_pet, :age_pet, :breed_pet, :nick_name_pet, :status_pet, :date_create_pet)`
	updatePetQuery = `UPDATE pets SET name_pet = :name_pet, 
										age_pet = :age_pet, 
										breed_pet = :breed_pet, 
										nick_name_pet = :nick_name_pet,
										date_update_pet = :date_update_pet 
						WHERE id = :id`
	deletePetQuery = `UPDATE PETS SET STATUS_PET = 0 WHERE id = $1`

	selectVaccinesQuery = `SELECT id, name_vaccine, date_given_vaccine, next_due_vaccine, id_pet, status_vaccine, date_create_vaccine, date_update_vaccine FROM vaccines`
	insertVaccineQuery  = `INSERT INTO vaccines (name_vaccine, date_given_vaccine, next_due_vaccine, id_pet, status_vaccine, date_create_vaccine) 
						VALUES (:name_vaccine, :date_given_vaccine, :next_due_vaccine, :id_pet, :status_vaccine, :date_create_vaccine)`
)
