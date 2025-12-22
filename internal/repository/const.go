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
)
